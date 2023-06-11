package main

import "fmt"

func main() {
	var numprowhile, variavel, par, impar, positivo, negativo int
	numprowhile = 5

	for numprowhile > 0 {
		fmt.Scan(&variavel)

		if variavel%2 == 0 {
			par++
		}
		if variavel%2 != 0 {
			impar++
		}
		if variavel < 0 {
			negativo++
		}
		if variavel > 0 {
			positivo++
		}

		numprowhile--
	}

	fmt.Printf("%d valor(es) par(es)\n", par)
	fmt.Printf("%d valor(es) impar(es)\n", impar)
	fmt.Printf("%d valor(es) positivo(s)\n", positivo)
	fmt.Printf("%d valor(es) negativo(s)\n", negativo)
}
