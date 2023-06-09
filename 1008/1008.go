package main

import (
	"fmt"
)

func main() {
	var matricula int
	var horas int
	var valor float32

	fmt.Scan(&matricula)
	fmt.Scan(&horas)
	fmt.Scan(&valor)

	fmt.Println("NUMBER =", matricula)
	fmt.Printf("SALARY = U$ %.2f\n", float32(horas)*valor)
}
