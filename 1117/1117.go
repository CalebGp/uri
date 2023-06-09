package main

import "fmt"

func main() {
	var nota, media, aux float32
	var validas int

	for validas != 2 {
		fmt.Scan(&nota)
		if nota >= 0 && nota <= 10 {
			validas++
			aux += nota
		} else {
			fmt.Println("nota invalida")
		}
	}

	media = aux / 2
	fmt.Printf("media = %.2f\n", media)
}
