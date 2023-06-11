package main

import "fmt"

func main() {
	var nome string
	var numerodenewtons, numerodetestes int

	fmt.Scan(&numerodetestes)

	for numerodetestes > 0 {
		fmt.Scan(&nome, &numerodenewtons)

		if nome == "Thor" {
			fmt.Println("Y")
		} else {
			fmt.Println("N")
		}

		numerodetestes--
	}
}
