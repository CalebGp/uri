package main

import (
	"fmt"
)

func main() {
	var id, t float64
	var qnt float64
	for {
		fmt.Scanf("%f", &id)
		if id < 0 {
			break
		} else {
			t += id
			qnt++
		}
	}
	fmt.Printf("%.2f\n", t/qnt)
}
