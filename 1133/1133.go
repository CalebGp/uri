package main

import (
	"fmt"
)

func main() {
	var a, b, aux int
	fmt.Scanf("%d\n%d", &a, &b)
	if a > b {
		aux = b
		b = a
		a = aux
	}
	for i := a + 1; i < b; i++ {
		if i%5 == 3 || i%5 == 2 {
			fmt.Printf("%d\n", i)
		}
	}
}
