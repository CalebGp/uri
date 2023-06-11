package main

import "fmt"

func main() {
	var numdetestes, num, in, out int
	fmt.Scan(&numdetestes)

	for i := 1; i <= numdetestes; i++ {
		fmt.Scan(&num)
		if num >= 10 && num <= 20 {
			in++
		} else {
			out++
		}
	}

	fmt.Printf("%d in\n", in)
	fmt.Printf("%d out\n", out)
}
