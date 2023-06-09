package main

import "fmt"

func main() {
	var n int //n é o valor total
	fmt.Scanf("%d", &n)
	var b int //b é o dinheiro restante
	fmt.Printf("%d\n", n)
	fmt.Printf("%d nota(s) de R$ 100,00\n", n/100)
	b = n % 100
	fmt.Printf("%d nota(s) de R$ 50,00\n", b/50)
	b = b % 50
	fmt.Printf("%d nota(s) de R$ 20,00\n", b/20)
	b = b % 20
	fmt.Printf("%d nota(s) de R$ 10,00\n", b/10)
	b = b % 10
	fmt.Printf("%d nota(s) de R$ 5,00\n", b/5)
	b = b % 5
	fmt.Printf("%d nota(s) de R$ 2,00\n", b/2)
	b = b % 2
	fmt.Printf("%d nota(s) de R$ 1,00\n", b/1)
}
