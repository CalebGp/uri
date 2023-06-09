package main

import (
	"fmt"
)

func main() {
	var media, mediafinal, notaex, a, b, c, d float64
	exame := false
	fmt.Scanf("%f %f %f %f", &a, &b, &c, &d)
	media = ((2 * a) + (3 * b) + (4 * c) + d) / 10
	fmt.Printf("Media: %.1f\n", media)
	if media >= 7.0 {
		fmt.Printf("Aluno aprovado.\n")
	} else if media < 5.0 {
		fmt.Printf("Aluno reprovado.\n")
	} else {
		exame = true
		fmt.Printf("Aluno em exame.\n")
	}
	if exame {
		fmt.Scanf("%f", &notaex)
		fmt.Printf("Nota do exame: %.1f\n", notaex)
		mediafinal = (media + notaex) / 2
		if mediafinal >= 5.0 {
			fmt.Printf("Aluno aprovado.\n")
		} else {
			fmt.Printf("Aluno reprovado.\n")
		}
		fmt.Printf("Media final: %.1f\n", mediafinal)
	}
}
