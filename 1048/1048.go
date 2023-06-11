package main

import "fmt"

var salario float32
var nvsalario float32
var reajuste float32
var porcentagem int

func funcao() {
	if salario <= 400 {
		nvsalario = salario * 1.15
		porcentagem = 15
	} else if salario > 400 && salario <= 800 {
		nvsalario = salario * 1.12
		porcentagem = 12
	} else if salario > 800 && salario <= 1200 {
		nvsalario = salario * 1.10
		porcentagem = 10
	} else if salario > 1200 && salario <= 2000 {
		nvsalario = salario * 1.07
		porcentagem = 7
	} else if salario > 2000 {
		nvsalario = salario * 1.04
		porcentagem = 4
	}
}

func main() {
	fmt.Scan(&salario)
	funcao()
	reajuste = nvsalario - salario
	fmt.Printf("Novo salario: %.2f\n", nvsalario)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Printf("Em percentual: %d %%\n", porcentagem)
}
