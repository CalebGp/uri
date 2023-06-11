package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println("O JOGO DUROU 24 HORA(S)")
	} else if a < b {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", b-a)
	} else if a > b {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", 24-(a-b))
	}
}
