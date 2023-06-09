package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for N != 0 && N < 1000 {
		if N >= 500 {
			if N >= 900 && N < 1000 {
				fmt.Printf("CM")
				N = N - 900
			} else {
				fmt.Printf("D")
				N = N - 500
			}
		} else if N < 500 && N >= 100 {
			if N >= 400 && N < 500 {
				fmt.Printf("CD")
				// fmt.Printf("M");
				N = N - 400
			} else {
				fmt.Printf("C")
				N = N - 100
			}
		} else if N < 100 && N >= 50 {
			if N >= 90 && N < 100 {
				fmt.Printf("XC")
				N = N - 90
			} else {
				fmt.Printf("L")
				N = N - 50
			}
		} else if N < 50 && N >= 10 {
			if N >= 40 {
				fmt.Printf("XL")
				N = N - 40
			} else {
				fmt.Printf("X")
				N = N - 10
			}
		} else if N < 10 && N >= 5 {
			if N == 9 {
				fmt.Printf("IX")
				N = N - 9
			} else {
				fmt.Printf("V")
				N = N - 5
			}
		} else {
			if N >= 4 && N < 5 {
				fmt.Printf("IV")
				N = N - 4
			} else {
				fmt.Printf("I")
				N = N - 1
			}
		}
	}
	fmt.Printf("\n")
}
