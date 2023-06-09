package main

import "fmt"

func main() {
	var salario float32
	fmt.Scan(&salario)

	if salario <= 2000.00 {
		fmt.Println("Isento")
	} else {
		var impostoderenda float32
		if salario > 2000.01 && salario <= 3000.00 {
			impostoderenda = (salario - 2000.01) * 0.08
		} else if salario > 3000.01 && salario <= 4500.00 {
			impostoderenda = 1000 * 0.08
			impostoderenda += (salario - 3000.01) * 0.18
		} else if salario > 4500.00 {
			impostoderenda = 1000 * 0.08
			impostoderenda += 1500 * 0.18
			impostoderenda += (salario - 4500.01) * 0.28
		}
		fmt.Printf("R$ %.2f\n", impostoderenda)
	}
}
