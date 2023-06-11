package main

import "fmt"

func main() {
	var numero int
	for numero != 2002 {
		fmt.Scan(&numero)
		if numero == 2002 {
			fmt.Println("Acesso Permitido")
		} else {
			fmt.Println("Senha Invalida")
		}
	}
}
