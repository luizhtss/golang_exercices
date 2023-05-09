package main

import (
	"fmt"
	"os"
)

// Definição das constantes
const SCORE = 350
const TAXA_DE_JUROS = 0.2

func valor_mensal_da_parcela(valor_emprestimo float64, tempo_emprestimo int) float64 {
	return (valor_emprestimo*TAXA_DE_JUROS)/float64(tempo_emprestimo) + valor_emprestimo/float64(tempo_emprestimo)
}

/*
O custo efetivo do empréstimo é calculado multiplicando-se o valor mensal de
cada parcela pelo tempo do empréstimo e subtraindo-o do valor requisitado pelo aplicante
*/
func custo_efetivo(valor_parcela float64, tempo_emprestimo int, valor_emprestimo float64) float64 {
	return (valor_parcela * float64(tempo_emprestimo)) - valor_emprestimo
}

/*
O empréstimo é aprovado se o aplicante possui renda maior que o valor mensal
da parcela, possui score de crédito considerado bom e possui percentual
de sua renda a ser comprometido compatível com o limite definido pelo score de crédito
*/
func situacao_emprestimo(renda float64, valor_parcela float64) string {
	if renda > valor_parcela && SCORE >= 501 && TAXA_DE_JUROS == 0.15 {
		return "APROVADO"
	}

	return "RECUSADO"
}

/*
Imprime todas as informações de análise de crédito do usuário e situação
do empréstimo
*/
func print_analise_de_credito(renda float64, valor_emprestimo float64, tempo_emprestimo int) {
	valor_mensal := valor_mensal_da_parcela(valor_emprestimo, tempo_emprestimo)
	custo_efetivo_emprestimo := custo_efetivo(valor_mensal, tempo_emprestimo, valor_emprestimo)

	fmt.Println("Análise de crédito para empréstimo")
	fmt.Println("----------------------------------")
	fmt.Println("Score de crédito:          ", SCORE)
	fmt.Printf("Renda:                      %.2f\n", renda)
	fmt.Printf("Valor do empréstimo:        %.2f\n", valor_emprestimo)
	fmt.Println("Tempo do empréstimo:       ", tempo_emprestimo)
	fmt.Printf("Valor mensal de parcela:    %.2f\n", valor_mensal)
	fmt.Printf("Taxa de juros:              %.2f%s\n", TAXA_DE_JUROS*100, "%")
	fmt.Printf("Custo efetivo:              %.2f\n", custo_efetivo_emprestimo)
	fmt.Println("Situação do empréstimo:    ", situacao_emprestimo(renda, valor_mensal))
}

/*
Um empréstimo é válido quanto otempo do empréstimo é divisível por 12. Essa
função verifica a validade do tempo_emprestimo.
*/
func eh_valido_tempo_emprestimo(tempo_emprestimo int) bool {
	return tempo_emprestimo%12 != 0
}

func main() {
	var renda float64 = 1000
	var valor_emprestimo float64 = 10000
	var tempo_emprestimo int = 12

	// Necessário validar se o tempo do empréstimo antes de executar.
	if eh_valido_tempo_emprestimo(tempo_emprestimo) {
		fmt.Fprint(os.Stderr, "Tempo do emprestimo não é divisível por 12.\n")
		os.Exit(1)
	}

	print_analise_de_credito(renda, valor_emprestimo, tempo_emprestimo)
}
