package main

import "fmt"

func main() {
	var tempo int
	fmt.Scanf("%d", &tempo)
	horas := tempo / 3600
	tempo = tempo - (horas * 3600)
	minutes := tempo / 60
	tempo = tempo - (minutes * 60)
	fmt.Printf("%d:%d:%d\n", horas, minutes, tempo)
}
