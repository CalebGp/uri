package main

import "fmt"

func main() {
	var x, y int
	for {
		fmt.Scan(&x)
		fmt.Scan(&y)
		if x == y {
			break
		}
		if x > y {
			fmt.Println("Decrescente")
		} else {
			fmt.Println("Crescente")
		}
	}
}
