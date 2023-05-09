package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Gera números aleatórios até encontrar um que seja divisível por 42.
func main() {
	rand.Seed(time.Now().Unix())
	contador := 0
	for {
		numero := rand.Int()
		fmt.Println(numero)
		contador++
		if numero%42 == 0 {
			break
		}

	}
	fmt.Printf("Fim após %d iterações.\n", contador)
}
