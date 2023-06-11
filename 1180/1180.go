package main

import "fmt"

func main() {
	var numdetestes int
	fmt.Scan(&numdetestes)
	vetor := make([]int, numdetestes)
	aux := 1000
	posicao := 0

	for i := 0; i < numdetestes; i++ {
		fmt.Scan(&vetor[i])
		if vetor[i] < aux {
			aux = vetor[i]
			posicao = i
		}
	}

	fmt.Printf("Menor valor: %d\n", aux)
	fmt.Printf("Posicao: %d\n", posicao)
}
