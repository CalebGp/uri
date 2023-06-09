package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	var menor, maior int
	if x < y {
		menor = x
		maior = y
	} else {
		menor = y
		maior = x
	}
	soma := 0
	for i := menor + 1; i < maior; i++ {
		if i%2 != 0 {
			soma += i
		}
	}

	fmt.Println(soma)
}
