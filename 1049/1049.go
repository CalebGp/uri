package main

import "fmt"

type Animal struct {
	Class map[string]map[string]string
}

var Animalia = map[string]Animal{
	"vertebrado": {
		Class: map[string]map[string]string{
			"ave": {
				"carnivoro": "aguia",
				"onivoro":   "pomba",
			},
			"mamifero": {
				"onivoro":   "homem",
				"herbivoro": "vaca",
			},
		},
	},
	"invertebrado": {
		Class: map[string]map[string]string{
			"inseto": {
				"hematofago": "pulga",
				"herbivoro":  "lagarta",
			},
			"anelideo": {
				"hematofago": "sanguessuga",
				"onivoro":    "minhoca",
			},
		},
	},
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	fmt.Println(Animalia[a].Class[b][c])
}
