package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Entry representa uma entrada em um mapa.
type Entry[K int | string, V constraints.Ordered] struct {
	key   K
	value V
}

// Map representa um mapa.
type Map[K int | string, V constraints.Ordered] struct {
	entries []Entry[K, V]
}

// getMapLength retorna o tamanho do mapa.
func getMapLength[K int | string, V constraints.Ordered](maps *Map[K, V]) int {
	return len(maps.entries)
}

// mapSize imprime o tamanho do mapa.
func mapSize[K int | string, V constraints.Ordered](maps *Map[K, V]) {
	fmt.Println(getMapLength(maps))
}

// findKeyIndex busca pelo índice de uma chave no mapa.
// Se a chave não existir, retorna -1.
func findKeyIndex[K int | string, V constraints.Ordered](maps *Map[K, V], key K) int {
	for index, entry := range maps.entries {
		if entry.key == key {
			return index
		}
	}
	return -1
}

// addEntry adiciona uma entrada em um mapa.
func addEntry[K int | string, V constraints.Ordered](maps *Map[K, V], key K, value V) {
	ikey := findKeyIndex(maps, key)
	// Verifica se a entry já existe para a chave.
	if ikey != -1 {
		maps.entries[ikey].value = value
	} else {
		// A entry não existe, adicione uma nova.
		maps.entries = append(maps.entries, Entry[K, V]{key, value})
	}
}

// getEntry imprime uma entrada em um mapa.
func getEntry[K int | string, V constraints.Ordered](maps *Map[K, V], key K) {
	// Verifique se a entry já existe para a chave.
	ikey := findKeyIndex(maps, key)
	if ikey != -1 {
		entry := maps.entries[ikey]
		fmt.Printf("[%v] \"%v\"\n", entry.key, entry.value)
	} else {
		fmt.Println("Chave não encontrada!")
	}
}

// printMap imprime o conteúdo de um mapa.
func printMap[K int | string, V constraints.Ordered](maps *Map[K, V]) {
	if getMapLength(maps) == 0 {
		fmt.Println("[]")
	} else {
		for _, entry := range maps.entries {
			fmt.Printf("[%v] \"%v\"\n", entry.key, entry.value)
		}
	}
	mapSize(maps)
}

// readInfinityLines lê input de múltiplas linhas do usuário.
// Permite ao usuário executar os comandos "print", "size", "add" e "exit".
func readInfinityLines(maps *Map[int, string]) {
	scanner := bufio.NewScanner(os.Stdin)
	var userInput []string
	var scannerInputText string

	fmt.Print("> ")
	for scanner.Scan() {
		scannerInputText = scanner.Text()
		userInput = strings.Split(strings.Join(strings.Fields(scannerInputText), " "), " ")
		if userInput[0] == "" {
			fmt.Print("> ")
			continue
		}

		switch userInput[0] {
		case "print":
			printMap(maps)
		case "size":
			fmt.Printf("%d\n", getMapLength(maps))
		case "add":
			cv_arg, err := strconv.Atoi(userInput[1])
			if err != nil {
				fmt.Println("Chave deve ser do tipo inteiro.")
			} else {
				addEntry(maps, cv_arg, userInput[2])
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
		}

		fmt.Print("> ")
	}
}

func main() {
	teste := Map[int, string]{}
	readInfinityLines(&teste)
}
