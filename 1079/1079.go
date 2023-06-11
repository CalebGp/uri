package main

import "fmt"

func main() {
	var a, b, c float64
	var media float64
	var numdetestes int

	fmt.Scan(&numdetestes)

	for i := 1; i <= numdetestes; i++ {
		fmt.Scan(&a, &b, &c)
		media = ((a * 2) + (b * 3) + (c * 5)) / 10
		fmt.Printf("%.1f\n", media)
	}
}
