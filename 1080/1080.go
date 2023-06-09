package main

import (
	"fmt"
)

func main() {
	var maior int
	posicao := 0
	var numeros [100]int
	for i := 0; i < 100; i++ {
		fmt.Scanf("%d\n", &numeros[i])
		if numeros[i] > maior {
			maior = numeros[i]
			posicao = i + 1
		}
	}
	fmt.Printf("%d\n", maior)
	fmt.Printf("%d\n", posicao)
}
