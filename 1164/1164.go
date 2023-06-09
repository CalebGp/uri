package main

import "fmt"

func main() {
	var numdetestes int
	fmt.Scan(&numdetestes)

	for i := 1; i <= numdetestes; i++ {
		var num, aux int
		fmt.Scan(&num)
		aux = 0
		for j := 1; j <= num/2; j++ {
			if num%j == 0 {
				aux += j
			}
		}
		if aux == num {
			fmt.Printf("%d eh perfeito\n", num)
		} else {
			fmt.Printf("%d nao eh perfeito\n", num)
		}
	}
}
