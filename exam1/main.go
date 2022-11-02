package main

import (
	"fmt"
)

func main() {
	rep := jusquauCarre(64)
	fmt.Println(rep)
}

func jusquauCarre(n int) int {
	reponse := 1
	var toCheck int
	estBon := true
	for estBon {
		toCheck = reponse * reponse
		if toCheck <= n {
			reponse++
		} else {
			estBon = false
		}
	}
	return reponse - 1
}
