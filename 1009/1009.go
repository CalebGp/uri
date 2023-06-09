package main

import "fmt"

func main() {
	var nome string
	var salario, totaldevendas float64
	fmt.Scanf("%s", &nome)
	fmt.Scanf("%f", &salario)
	fmt.Scanf("%f", &totaldevendas)
	fmt.Printf("TOTAL = R$ %.2f\n", salario+totaldevendas*0.15)
}
