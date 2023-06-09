package main

import "fmt"

func main() {
	var numdetestes, total, quantia, numdesapos, numdecoelhos, numderatos int
	var tipo string
	fmt.Scan(&numdetestes)

	for i := 1; i <= numdetestes; i++ {
		fmt.Scan(&quantia, &tipo)
		total += quantia
		if tipo == "C" {
			numdecoelhos += quantia
		}
		if tipo == "S" {
			numdesapos += quantia
		}
		if tipo == "R" {
			numderatos += quantia
		}
	}
	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Printf("Total de coelhos: %d\n", numdecoelhos)
	fmt.Printf("Total de ratos: %d\n", numderatos)
	fmt.Printf("Total de sapos: %d\n", numdesapos)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", (float64(numdecoelhos)/float64(total))*100.00)
	fmt.Printf("Percentual de ratos: %.2f %%\n", (float64(numderatos)/float64(total))*100.00)
	fmt.Printf("Percentual de sapos: %.2f %%\n", (float64(numdesapos)/float64(total))*100.00)
}
