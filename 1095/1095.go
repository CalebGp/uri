package main

import "fmt"

func main() {
	i := 60
	j := 1

	for i >= 0 {
		fmt.Printf("I=%d J=%d\n", j, i)
		i -= 5
		j += 3
	}
}
