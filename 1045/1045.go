package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	if a >= (b+c) || b >= (a+c) || c >= (a+b) {
		fmt.Println("NAO FORMA TRIANGULO")
	} else if a*a == (b*b+c*c) || b*b == (a*a+c*c) || c*c == (a*a+b*b) {
		fmt.Println("TRIANGULO RETANGULO")
	} else if a*a > (b*b+c*c) || b*b > (a*a+c*c) || c*c > (a*a+b*b) {
		fmt.Println("TRIANGULO OBTUSANGULO")
	} else if a*a < (b*b+c*c) || b*b < (a*a+c*c) || c*c < (a*a+b*b) {
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	if a == b && a == c {
		fmt.Println("TRIANGULO EQUILATERO")
	}

	if (a == b && a != c) || (a == c && a != b) || (b == c && b != a) {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}
