package main

import (
	"fmt"
)

func main() {
	var n int
	var vi, vg int
	for n != 2 {
		var gi, gg int
		fmt.Scanf("%d %d", &gi, &gg)
		if gi > gg {
			vi++
		} else if gg > gi {
			vg++
		} else {
			fmt.Printf("Nao houve vencedor\n")
		}
		fmt.Scanf("Novo grenal (1-sim 2-nao)\n")
		fmt.Println("Novo grenal (1-sim 2-nao)")
		var continuar int
		fmt.Scanln(&continuar)
		if continuar != 1 {
			break
		}
	}

}
