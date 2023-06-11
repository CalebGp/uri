package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 7; j > 4; j-- {
			fmt.Printf("I=%d J=%d\n", i, j)
		}
		i++
	}
}
