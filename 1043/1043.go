package main

import "fmt"

func main() {
	var a, b, c, area, perimetro float64
	fmt.Scan(&a, &b, &c)

	perimetro = a + b + c
	area = (a + b) * c / 2

	if a < (b+c) && b < (a+c) && c < (a+b) {
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else {
		fmt.Printf("Area = %.1f\n", area)
	}
}
