package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scan(&num2)

	for num3 := 1; num3 <= num2; num3++ {
		fmt.Scan(&num1)

		if num1%2 == 0 {
			fmt.Println("0")
		} else {
			fmt.Println("1")
		}
	}
}
