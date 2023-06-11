package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		numaoquadrado := int(math.Pow(float64(i), 2))
		numaocubo := int(math.Pow(float64(i), 3))
		fmt.Printf("%d %d %d\n", i, numaoquadrado, numaocubo)
	}
}
