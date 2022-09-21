package main

import (
	"flag"
	"fmt"
	m "mod/map"
	"strconv"
	"strings"
)

func main() {
	fun := flag.String("-c", "", "description")
	fun1 := flag.String("-n", "", "description")
	flag.Parse()
	s := strings.Split(*fun1, ",")
	ints := make([]int, len(s))
	for i, v := range s {
		ints[i], _ = strconv.Atoi(v)
	}
	switch *fun {
	case "s":
		fmt.Println(m.Somme(ints))

	case "m":
		fmt.Println(m.Moyenne(ints))

	case "sd":
		fmt.Println(m.StandardDeviation(ints))

	case "max":
		fmt.Println(m.Max(ints))

	case "min":
		fmt.Println(m.Min(ints))

	case "med":
		fmt.Println(m.Median(ints))

	default:
		fmt.Println("error")
	}
}
