package main

import "fmt"

func main() {
	var km, velocidade float64
	var distancia float64
	fmt.Scanf("%f\n %f", &km, &velocidade)
	distancia = km * velocidade
	fmt.Printf("%.3f\n", distancia/12)
}
