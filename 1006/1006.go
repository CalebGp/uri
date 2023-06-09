package main

import "fmt"

func main() {
	var A, B, C, D, dif int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &D)
	dif = (A * B) - (C * D)
	fmt.Printf("DIFERENCA = %d\n", dif)
}
