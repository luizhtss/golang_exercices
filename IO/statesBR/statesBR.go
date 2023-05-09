package main

import "fmt"

type Estado struct {
	nome      string
	sigla     string
	capital   string
	populacao int
	regiao    string
}

func main() {
	// inicializa o mapa de Estados com algumas informações
	estados := map[string]Estado{
		"Paraná":    {"Paraná", "PR", "Curitiba", 11433957, "Sul"},
		"São Paulo": {"São Paulo", "SP", "São Paulo", 46649132, "Sudeste"},
		"Roraima":   {"Roraima", "RR", "Boa Vista", 605761, "Norte"},
	}

	// insere o Rio Grande do Norte no map.
	estados["Rio Grande do Norte"] = Estado{"Rio Grande do Norte", "RN", "Natal", 3534165, "Nordeste"}

	// realiza algumas consultas no mapa de Estados
	fmt.Println("Qual a capital de Paraná?")
	fmt.Println("-", estados["Paraná"].capital)

	fmt.Println("Qual a população de São Paulo?")
	fmt.Println("-", estados["São Paulo"].populacao, "habitantes")

	fmt.Println("Qual a sigla de Roraima?")
	fmt.Println("-", estados["Roraima"].sigla)

	fmt.Println("Em que região Rio Grande do Norte está localizado?")
	fmt.Println("-", estados["Rio Grande do Norte"].regiao)
}
