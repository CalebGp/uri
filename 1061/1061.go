package main

import "fmt"

func main() {
	var dia, diafim, hora, horafim, minuto, minutofinal, segundo, segundofinal int

	fmt.Scanf("Dia %d", &dia)
	fmt.Scanf("%d : %d : %d\n", &hora, &minuto, &segundo)
	fmt.Scanf("Dia %d", &diafim)
	fmt.Scanf("%d : %d : %d", &horafim, &minutofinal, &segundofinal)

	segundo = segundofinal - segundo
	minuto = minutofinal - minuto
	hora = horafim - hora
	dia = diafim - dia

	if segundo < 0 {
		segundo += 60
		minuto--
	}

	if minuto < 0 {
		minuto += 60
		hora--
	}

	if hora < 0 {
		hora += 24
		dia--
	}

	fmt.Printf("%d dia(s)\n", dia)
	fmt.Printf("%d hora(s)\n", hora)
	fmt.Printf("%d minuto(s)\n", minuto)
	fmt.Printf("%d segundo(s)\n", segundo)
}
