package main

import (
	"fmt"
)

func main() {
	var horaInicial, minutoInicial, horaFinal, minutoFinal int

	fmt.Scan(&horaInicial, &minutoInicial, &horaFinal, &minutoFinal)

	minutosInicial := horaInicial*60 + minutoInicial
	minutosFinal := horaFinal*60 + minutoFinal

	duracaoMinutos := minutosFinal - minutosInicial

	if duracaoMinutos <= 0 {
		duracaoMinutos += 24 * 60
	}

	horas := duracaoMinutos / 60
	minutos := duracaoMinutos % 60

	// Exibir o resultado
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", horas, minutos)
}
