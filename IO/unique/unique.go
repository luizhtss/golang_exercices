package main

import "fmt"

// Verifica se um determinado inteiro está contido em uma slice de inteiros.
func existsOnSlice(slice []int, find int) bool {
	for _, value := range slice {
		if value == find {
			return true
		}
	}
	return false
}

func main() {
	// Cria uma array de slice que contém casos de teste.
	values := [3][]int{
		{98, 76, 68, 76, 76, 48, 73, 16, 16, 99},
		{1, 1, 1, 1, 1, 1},
		{22, 13, 13, 12, 5, 5, 45, 16, 11, 77},
	}
	// Percorre todos os casos de teste.
	for _, slices := range values {
		unique_slice := []int{}
		// Itera por cada elemento da slice e adiciona o elemento à slice unique_slice apenas se o elemento ainda não existe nela.
		for _, value := range slices {
			if !existsOnSlice(unique_slice, value) {
				unique_slice = append(unique_slice, value)
			}
		}
		// Imprime a nova slice, agora sem elementos repetidos.
		fmt.Println(unique_slice)
	}

}
