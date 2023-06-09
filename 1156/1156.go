package main

import (
	"fmt"
)

func main() {
	var s float64
	s = 1
	num := 3.00
	for i := 2.00; num <= 39; i *= 2 {
		s += num / i
		num += 2
	}
	fmt.Printf("%.2f\n", s)
}
