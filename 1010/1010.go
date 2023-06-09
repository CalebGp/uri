package main

import "fmt"

func main() {
	var c1, c2 int
	var n1, n2, p1, p2 float64
	fmt.Scanf("%d %f %f", &c1, &n1, &p1)
	fmt.Scanf("%d %f %f", &c2, &n2, &p2)
	total := (n1 * p1) + (n2 * p2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
