package main

import (
	"fmt"
	"runtime"
	"time"
)

var i, j int = 0, 1

func main() {
	i = 5
	var nombre = 15
	brombre := 10
	fmt.Println(i, nombre, brombre)

	var t []int = []int{1, 2, 3}
	b := [3]int{1, 2, 3}
	s := b[1:3]
	t[2] = 4
	fmt.Println(t, b, s)

	if calculus := 7 % 2; calculus == 0 {
		fmt.Println("ok")
	}

	switch os := runtime.GOOS; os{
	case "darwom":
			fmt.Println("OS X.-")
	case "linux":
			fmt.Println("Linux")
	default:
			fmt.Println("%s.\n", os)
	}

	v := time.Now()
	switch {
	case v.Hour() < 12:
			fmt.Println("Avant-midi")
	default:
			fmt.Println("Apres-diner")
	}

	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}

	j:= 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	
	l := []string{"patate", "carotte", "Michel"}
	for _, v := range l{
		fmt.Println(v)
	}
	a, m := ajouterEtMultiplier(1,2)
	fmt.Println(a)
	fmt.Println(m)

}

func ajouterEtMultiplier(x int, y int) (int, int) {
	return x + y, x * y
}
