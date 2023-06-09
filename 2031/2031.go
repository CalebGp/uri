package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	var jogador1, jogador2 string
	for i := 1; i <= N; i++ {
		fmt.Scanf("%s\n%s\n", &jogador1, &jogador2)
		if jogador1 == "pedra" && jogador2 == "ataque" {
			fmt.Printf("Jogador 2 venceu\n")
		}
		if jogador1 == "ataque" && jogador2 == "pedra" {
			fmt.Printf("Jogador 1 venceu\n")
		}
		if jogador1 == "pedra" && jogador2 == "papel" {
			fmt.Printf("Jogador 1 venceu\n")
		}
		if jogador1 == "papel" && jogador2 == "pedra" {
			fmt.Printf("Jogador 2 venceu\n")
		}
		if jogador1 == "ataque" && jogador2 == "papel" {
			fmt.Printf("Jogador 1 venceu\n")
		}
		if jogador1 == "papel" && jogador2 == "ataque" {
			fmt.Printf("Jogador 2 venceu\n")
		}
		if jogador1 == "papel" && jogador2 == "papel" {
			fmt.Printf("Ambos venceram\n")
		}
		if jogador1 == "pedra" && jogador2 == "pedra" {
			fmt.Printf("Sem ganhador\n")
		}
		if jogador1 == "ataque" && jogador2 == "ataque" {
			fmt.Printf("Aniquilacao mutua\n")
		}
	}
}
