package main

import "fmt"

func main() {
	var jogo [6]int
	var sorteados [6]int

	for i := 0; i < 6; i++ {
		fmt.Scan(&jogo[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scan(&sorteados[i])
	}

	acertos := 0

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if jogo[i] == sorteados[j] {
				acertos++
			}
		}
	}

	switch acertos {
	case 0, 1, 2:
		fmt.Println("azar")
	case 3:
		fmt.Println("terno")
	case 4:
		fmt.Println("quadra")
	case 5:
		fmt.Println("quina")
	case 6:
		fmt.Println("sena")
	}
}
