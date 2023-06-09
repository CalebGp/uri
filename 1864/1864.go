package main

import "fmt"

func main() {
	var numdeletras int
	fmt.Scan(&numdeletras)

	mensagem := "LIFE IS NOT A PROBLEM TO BE SOLVED"
	for i := 0; i < numdeletras; i++ {
		fmt.Printf("%c", mensagem[i])
	}
	fmt.Println()
}
