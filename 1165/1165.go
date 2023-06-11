package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var testes int
	fmt.Scan(&testes)

	for testes > 0 {
		var num int
		fmt.Scan(&num)

		if isPrime(num) {
			fmt.Printf("%d eh primo\n", num)
		} else {
			fmt.Printf("%d nao eh primo\n", num)
		}

		testes--
	}
}
