package main

import "fmt"

func main() {
	var acertos int
	var CorrectAnswer uint8
	fmt.Scanf("%d\n", &CorrectAnswer)
	var concorrentes = [5]uint8{}
	fmt.Scanf("%d %d %d %d %d", &concorrentes[0], &concorrentes[1], &concorrentes[2], &concorrentes[3], &concorrentes[4])
	for i := 0; i < 5; i++ {
		if concorrentes[i] == CorrectAnswer {
			acertos++
		}
	}
	fmt.Println(acertos)
}