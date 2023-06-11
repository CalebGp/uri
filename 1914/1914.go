package main

import "fmt"

func main() {
	var plv [4]string
	var numtt, num1, num2, soma int

	fmt.Scan(&numtt)

	for numtt > 0 {
		fmt.Scan(&plv[0], &plv[1], &plv[2], &plv[3])
		fmt.Scan(&num1, &num2)

		soma = num1 + num2

		if soma%2 == 0 {
			if plv[1] == "PAR" {
				fmt.Println(plv[0])
			} else {
				fmt.Println(plv[2])
			}
		} else {
			if plv[1] == "IMPAR" {
				fmt.Println(plv[0])
			} else {
				fmt.Println(plv[2])
			}
		}

		numtt--
	}
}
