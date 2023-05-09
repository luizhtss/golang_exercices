package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type EstatisticasAno struct {
	maiorNumeroNascidos   int
	menorNumeroNascidos   int
	totalNumeroNascidos   int
	nascimentosMunicipios []int
	mediaNascidos         int
	desvioPadraoNascidos  float64
}

type EstatisticasMunicipio struct {
	codigo      int
	municipio   string
	nascimentos map[int]int // K (ano), V (nascimentos).
}

type TaxaCrescrimentoMunicipio struct {
	nome            string
	taxaCrescimento float64
}

func (e EstatisticasAno) calcularDesvioPadraoNascidos() float64 {
	var soma float64
	for _, nascimentos := range e.nascimentosMunicipios {
		soma += math.Pow(float64(nascimentos-e.mediaNascidos), 2.0)
	}
	return math.Sqrt(float64(1.0/float64(len(e.nascimentosMunicipios))) * soma)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso correto: nascimentos <caminho_do_arquivo>")
		return
	}

	caminhoCSV := os.Args[1]
	dataAno := map[int]EstatisticasAno{}
	dataCidade := map[int]EstatisticasMunicipio{}
	gerarEstatisticas(&dataAno, &dataCidade, caminhoCSV)
	gerarArquivosEstatisticas(dataAno)
	plotarHistograma()
	gerarArquivoNascimentosAlvos(dataCidade)
	gerarGraficoLinha()
	imprimeMenorEMaiorTaxaDeCrescimento(dataCidade)
}

func imprimeMenorEMaiorTaxaDeCrescimento(dataCidade map[int]EstatisticasMunicipio) {
	// Cria as variáveis de maior e menor taxa
	var maiorTaxa, menorTaxa TaxaCrescrimentoMunicipio
	maiorTaxa.taxaCrescimento = math.Inf(-1)
	menorTaxa.taxaCrescimento = math.Inf(1)

	// Varre cada cidade para olhar os números de nascidos
	for _, estatisticas := range dataCidade {
		// Calcula a taxa de crescimento caso exista os anos de 2016 e 2020
		// e retorna se deu certo na variável 'ok'
		tc, ok := taxaDeCrescimentoRelativa(estatisticas.nascimentos)
		if ok {
			if maiorTaxa.taxaCrescimento < tc {
				maiorTaxa.taxaCrescimento = tc
				maiorTaxa.nome = estatisticas.municipio
			}

			if menorTaxa.taxaCrescimento > tc {
				menorTaxa.taxaCrescimento = tc
				menorTaxa.nome = estatisticas.municipio
			}
		}
	}

	// Calcula a porcentagem de c1rescimento de ambos os municípios selecionados
	// com a maior e menor taxa
	maiorTaxa.taxaCrescimento = porcentagemDeCrescimento(maiorTaxa.taxaCrescimento)
	menorTaxa.taxaCrescimento = porcentagemDeCrescimento(menorTaxa.taxaCrescimento)

	// Caso alguma taxa seja menor do que zero, deve ser informada como 'taxa de queda'
	tipoVariacaoMaiorTaxa, tipoVariacaoMenorTaxa := "taxa de crescimento", "taxa de crescimento"
	if maiorTaxa.taxaCrescimento < 0 {
		tipoVariacaoMaiorTaxa = "taxa de queda"
	}
	if menorTaxa.taxaCrescimento < 0 {
		tipoVariacaoMenorTaxa = "taxa de queda"
	}

	// Imprime os municípios com maior e menor taxa
	fmt.Printf("Município com maior %s 2016-2020: %s (%.2f%%)\n", tipoVariacaoMaiorTaxa, maiorTaxa.nome, maiorTaxa.taxaCrescimento)
	fmt.Printf("Município com maior %s 2016-2020: %s (%.2f%%)\n", tipoVariacaoMenorTaxa, menorTaxa.nome, menorTaxa.taxaCrescimento)
}

func taxaDeCrescimentoRelativa(m map[int]int) (float64, bool) {
	nascimentos_2020, ano_2020_existe := m[2020]
	nascimentos_2016, ano_2016_existe := m[2016]
	if ano_2020_existe && ano_2016_existe {
		return float64(nascimentos_2020) / float64(nascimentos_2016), true
	}

	return 0, false
}

func porcentagemDeCrescimento(tc float64) float64 {
	return 100 * (tc - 1)
}

func gerarEstatisticas(dataAno *map[int]EstatisticasAno, dataCidade *map[int]EstatisticasMunicipio, filePath string) {
	// Abre o arquivo CSV
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Cria o leitor CSV
	reader := csv.NewReader(file)

	// Lê todas as linhas do arquivo
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	// Percorre as colunas do arquivo
	for j := 1; j < len(lines[0])-1; j++ {
		// Obtém o ano a partir do índice da coluna
		ano := 1993 + j
		// Inicializa a estatística para o ano atual
		tempAno := EstatisticasAno{menorNumeroNascidos: 999999, maiorNumeroNascidos: 0}
		registrosValidos := 0
		// Percorre as linhas do arquivo.
		for i := 1; i < len(lines)-1; i++ {
			// Obtém o total do ano.
			registroAno, _ := strconv.Atoi(lines[i][j])

			// Extrai o código e o nome do município.
			split := strings.Split(lines[i][0], " ")
			municipio := strings.Join(split[1:], " ")
			codigo, _ := strconv.Atoi(split[0])

			// Ignora campo sem registro.
			if registroAno == 0 {
				continue
			}
			// Verifica se o código se encontra no map da cidade.
			if _, existe := (*dataCidade)[codigo]; existe {
				// A chave existe no map
				(*dataCidade)[codigo].nascimentos[ano] += registroAno
			} else {
				// Não existe, então adicione.
				(*dataCidade)[codigo] = EstatisticasMunicipio{
					codigo:    codigo,
					municipio: municipio,
					nascimentos: map[int]int{
						ano: registroAno,
					},
				}
			}
			registrosValidos++
			// Adiciona o total daquele ano na propriedade correspondente da estrutura
			tempAno.totalNumeroNascidos += registroAno

			// Adiciona o menor número de nascimento registrados.
			if registroAno < tempAno.menorNumeroNascidos {
				tempAno.menorNumeroNascidos = registroAno
			}
			// Adiciona o maior número de nascimento registrados.
			if registroAno > tempAno.maiorNumeroNascidos {
				tempAno.maiorNumeroNascidos = registroAno
			}
			// Adiciona o número de nascimento registrados na slice.
			tempAno.nascimentosMunicipios = append(tempAno.nascimentosMunicipios, registroAno)

			// Verificação de integridade da soma dos registrados.
			if i == len(lines)-2 {
				registroCSV, _ := strconv.Atoi(lines[i+1][j])
				if registroCSV != tempAno.totalNumeroNascidos {
					println("[ERRO] A soma total de registrados não corresponde com o valor registrado no CSV.")
					os.Exit(1)
				}
			}
		}
		tempAno.mediaNascidos = tempAno.totalNumeroNascidos / registrosValidos
		tempAno.desvioPadraoNascidos = tempAno.calcularDesvioPadraoNascidos()
		(*dataAno)[ano] = tempAno
	}
}

func gerarArquivosEstatisticas(data map[int]EstatisticasAno) {
	// Cria o arquivo estatisticas.csv
	estatisticasFile, err := os.Create("estatisticas.csv")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo estatisticas.csv:", err)
		return
	}
	defer estatisticasFile.Close()

	// Cria o escritor CSV
	estatisticasWriter := csv.NewWriter(estatisticasFile)
	cabecalho := []string{"Ano", "maiorNumeroNascimentos", "menorNumeroNascimentos", "mediaNumeroNascimentos", "desvioNumeroNascimento", "totalNascimentos"}
	estatisticasWriter.Write(cabecalho)

	// Percorre os dados para escrever as linhas do arquivo estatisticas.csv
	for ano, estatisticas := range data {
		linha := []string{strconv.Itoa(ano), strconv.Itoa(estatisticas.maiorNumeroNascidos), strconv.Itoa(estatisticas.menorNumeroNascidos), strconv.Itoa(estatisticas.mediaNascidos), strconv.FormatFloat(estatisticas.desvioPadraoNascidos, 'f', 2, 64), strconv.Itoa(estatisticas.totalNumeroNascidos)}
		estatisticasWriter.Write(linha)
	}
	estatisticasWriter.Flush()
	println("> Arquivo estatisticas.csv gerado.")
	// Cria o arquivo totais.dat
	totaisFile, err := os.Create("totais.dat")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo totais.dat:", err)
		return
	}
	defer totaisFile.Close()

	// Escreve no arquivo totais.dat o número total de nascimentos em cada ano
	for ano, estatisticas := range data {
		linha := fmt.Sprintf("%d %d\n", ano, estatisticas.totalNumeroNascidos)
		totaisFile.WriteString(linha)
	}
	println("> Arquivo totais.dat gerado.")
}

func plotarHistograma() {
	// Prepara o comando que gera o histograma.
	cmd := exec.Command("gnuplot", "-e", "filename='totais.dat'", "histograma.gnuplot")

	// Executa o comando que gera o histograma.
	err := cmd.Run()

	// Verifica se o comando foi executado com sucesso.
	if err != nil {
		fmt.Println("[ERRO] Erro ao executar o comando gnuplot:", err)
		return
	}
	fmt.Println("[INFO] Histograma gerado com sucesso.")
}

func lerArquivoAlvos() []int {
	println("> Lendo arquivo alvos.dat")
	// Abre o arquivo alvos.dat
	file, err := os.Open("alvos.dat")
	if err != nil {
		println("[ERRO] Falha ao ler o arquivo alvos.dat: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// Cria o leitor CSV
	reader := csv.NewReader(file)

	// Lê todas as linhas do arquivo
	lines, err := reader.ReadAll()
	if err != nil {
		println("[ERRO] Falha ao ler o arquivo alvos.dat: ", err)
		os.Exit(1)
	}

	// Extrai os códigos dos municípios alvo
	alvos := []int{}
	for _, line := range lines {
		codigo, err := strconv.Atoi(line[0])
		if err != nil {
			println("[ERRO] Falha ao ler o arquivo alvos.dat: ", err)
			os.Exit(1)
		}
		alvos = append(alvos, codigo)
	}

	return alvos
}

func gerarArquivoNascimentosAlvos(dataCidade map[int]EstatisticasMunicipio) {
	// Cria o arquivo nascimentos-alvos.dat
	file, err := os.Create("nascimentos-alvos.dat")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo nascimentos-alvos.dat:", err)
		return
	}
	defer file.Close()
	alvos := lerArquivoAlvos()
	fmt.Printf("Municípios definidos como alvo (%d):\n", len(alvos))
	// Escreve o cabeçalho no arquivo com os nomes dos municípios alvo
	header := "Ano"
	for _, codigo := range alvos {
		municipio := dataCidade[codigo].municipio
		println(municipio)
		header += "," + municipio
	}
	file.WriteString(header + "\n")

	// Percorre os anos e escreve as linhas correspondentes aos nascimentos dos municípios alvo
	for ano := 1994; ano < 2021; ano++ {
		linha := strconv.Itoa(ano)

		for _, codigo := range alvos {
			nascimentos := dataCidade[codigo].nascimentos[ano]
			linha += "," + strconv.Itoa(nascimentos)
		}

		file.WriteString(linha + "\n")
	}

	fmt.Println("> Arquivo nascimentos-alvos.dat gerado.")
}

func gerarGraficoLinha() {
	// Prepara o comando que gera o gráfico de linha.
	cmd := exec.Command("gnuplot", "-e", "datafile='nascimentos-alvos.dat'", "linechart.gnuplot")

	// Executa o comando que gera o gráfico de linha.
	err := cmd.Run()

	// Verifica se o comando foi executado com sucesso.
	if err != nil {
		fmt.Println("[ERRO] Erro ao executar o comando gnuplot:", err)
		return
	}
	fmt.Println("[INFO] Gráfico de linha gerado com sucesso.")
}
