package main

import "fmt"

func main() {
	var T1, T2, T3, T4 int
	fmt.Scanf("%d %d %d %d", &T1, &T2, &T3, &T4)
	result := (T1 + T2 + T3 + T4) - 3
	fmt.Println(result)
}
