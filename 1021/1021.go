package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	inteiro := int(math.Round(n * 100.00))
	aux1 := inteiro

	fmt.Println("NOTAS:")
	fmt.Printf("%d nota(s) de R$ 100.00\n", aux1/10000)
	aux1 %= 10000
	fmt.Printf("%d nota(s) de R$ 50.00\n", aux1/5000)
	aux1 %= 5000
	fmt.Printf("%d nota(s) de R$ 20.00\n", aux1/2000)
	aux1 %= 2000
	fmt.Printf("%d nota(s) de R$ 10.00\n", aux1/1000)
	aux1 %= 1000
	fmt.Printf("%d nota(s) de R$ 5.00\n", aux1/500)
	aux1 %= 500
	fmt.Printf("%d nota(s) de R$ 2.00\n", aux1/200)
	aux1 %= 200

	fmt.Println("MOEDAS:")
	fmt.Printf("%d moeda(s) de R$ 1.00\n", aux1/100)
	aux1 %= 100
	fmt.Printf("%d moeda(s) de R$ 0.50\n", aux1/50)
	aux1 %= 50
	fmt.Printf("%d moeda(s) de R$ 0.25\n", aux1/25)
	aux1 %= 25
	fmt.Printf("%d moeda(s) de R$ 0.10\n", aux1/10)
	aux1 %= 10
	fmt.Printf("%d moeda(s) de R$ 0.05\n", aux1/5)
	aux1 %= 5
	fmt.Printf("%d moeda(s) de R$ 0.01\n", aux1/1)
}
