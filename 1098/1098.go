package main

import "fmt"

func main() {
	i := 0.0
	j := 1.0

	for i <= 2.1 {
		fmt.Printf("I=%.1f J=%.1f\n", i, j+i)
		fmt.Printf("I=%.1f J=%.1f\n", i, j+1+i)
		fmt.Printf("I=%.1f J=%.1f\n", i, j+2+i)
		i += 0.2
	}
}
