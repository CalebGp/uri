package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, maiorAb, maior float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	maiorAb = (a + b + math.Abs(a-b)) / 2
	maior = (maiorAb + c + math.Abs(maiorAb-c)) / 2
	fmt.Printf("%.0f eh o maior\n", maior)
}
