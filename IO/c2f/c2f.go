package main

import "fmt"

// Função que converte graus Celsius em graus Fahrenheit.
// Recebe um valor em Celsius como float64 e retorna um valor em Fahrenheit como float64.
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// Função que desenha uma tabela de conversão de temperatura.
// Recebe os valores Celsius e Fahrenheit como strings.
// A tabela é desenhada com bordas e os valores são formatados com uma largura de 5 caracteres.
func drawTable(valueCelsius string, valueFahrenheit string) {
	// Desenha cabeçalho da tabela, se o valor de Celsius for -40.0.
	if valueCelsius == "-40.0" {
		fmt.Println("=================")
		fmt.Println("|    ºC |    ºF |")
		fmt.Println("=================")
	}
	// Desenha linha da tabela com os valores de Celsius e Fahrenheit.
	fmt.Printf("| %5s | %5s |\n", valueCelsius, valueFahrenheit)
	// Desenha rodapé da tabela, se o valor de Celsius for 100.0.
	if valueCelsius == "100.0" {
		fmt.Println("=================")
	}
}

func main() {
	// Loop para converter valores de Celsius para Fahrenheit e desenhar a tabela de conversão.
	for celsius := -40.0; celsius <= 100.0; celsius += 5 {
		fahrenheit := celsiusToFahrenheit(celsius)
		string_celsius := fmt.Sprintf("%.1f", celsius)
		string_fahrenheit := fmt.Sprintf("%.1f", fahrenheit)
		drawTable(string_celsius, string_fahrenheit)
	}
}
