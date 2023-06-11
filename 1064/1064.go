package main

import "fmt"

func main() {
	var num float32
	var media float32
	positivo := 0
	somadosnumpositivos := float32(0)

	for numdetestes := 1; numdetestes <= 6; numdetestes++ {
		fmt.Scan(&num)
		if num > 0 {
			positivo++
			somadosnumpositivos += num
		}
	}

	media = somadosnumpositivos / float32(positivo)

	fmt.Printf("%d valores positivos\n", positivo)
	fmt.Printf("%.1f\n", media)
}
