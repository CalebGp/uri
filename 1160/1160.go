package main

import (
	"fmt"
	"math"
)

func main() {
	var numTests int
	var a, b int
	var taxaA, taxaB float64
	fmt.Scanf("%d\n", &numTests)
	for i := 1; i <= numTests; i++ {
		fmt.Scanf("%d %d %f %f\n", &a, &b, &taxaA, &taxaB)
		anos := 0
		for a <= b && anos <= 100 {
			a += int(math.Floor(float64(a) * taxaA / 100))
			b += int(math.Floor(float64(b) * taxaB / 100))
			anos++
		}
		if anos <= 100 {
			fmt.Printf("%d anos.\n", anos)
		} else {
			fmt.Printf("Mais de 1 seculo.\n")
		}
	}
}
