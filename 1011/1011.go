package main

import "fmt"

func main() {
	var raio float64
	pi := 3.14159
	fmt.Scanf("%f", &raio)
	total := (4.0 / 3.0) * pi * raio * raio * raio
	fmt.Printf("VOLUME = %.3f\n", total)
}
