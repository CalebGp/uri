package main

import "fmt"

func main() {
	alcool := 0
	diesel := 0
	gasolina := 0
	var num int
	fmt.Scanf("%d", &num)
	for num != 4 {
		fmt.Scanf("%d", &num)
		switch num {
		case 1:
			alcool++
			break
		case 2:
			gasolina++
			break
		case 3:
			diesel++
			break
		default:
			break
		}
	}

	fmt.Printf("MUITO OBRIGADO\n")
	fmt.Printf("Alcool: %d\n", alcool)
	fmt.Printf("Gasolina: %d\n", gasolina)
	fmt.Printf("Diesel: %d\n", diesel)
}
