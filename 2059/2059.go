package main

import "fmt"

func main() {
	var p, j1, j2, r, a int
	fmt.Scan(&p, &j1, &j2, &r, &a)

	soma := j1 + j2

	if r == 1 {
		if a == 0 {
			fmt.Println("Jogador 1 ganha!")
		} else if a == 1 {
			fmt.Println("Jogador 2 ganha!")
		}
	} else if soma%2 != p {
		fmt.Println("Jogador 1 ganha!")
	} else {
		fmt.Println("Jogador 2 ganha!")
	}
}
