package main

import (
	"fmt"
)

func main() {
	var alturaPulo, numCanos int
	fmt.Scan(&alturaPulo, &numCanos)

	cano := make([]int, numCanos)
	fmt.Scan(&cano[0])

	for i := 1; i < numCanos; i++ {
		fmt.Scan(&cano[i])

		if cano[i] >= cano[i-1] && cano[i]-cano[i-1] > alturaPulo {
			fmt.Println("GAME OVER")
			return
		} else if cano[i-1]-cano[i] > alturaPulo {
			fmt.Println("GAME OVER")
			return
		}
	}

	fmt.Println("YOU WIN")
}
