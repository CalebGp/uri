package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(num)
}
