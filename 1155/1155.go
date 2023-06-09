package main

import (
	"fmt"
)

func main() {
	var s float64
	s = 0
	for i := 1.00; i <= 100; i++ {
		s += 1.00 / i
	}
	fmt.Printf("%.2f\n", s)
}
