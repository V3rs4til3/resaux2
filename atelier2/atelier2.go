package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var cache = make(map[int]int)

func main() {
	var s string = "La patate douce n'est pas apparentée à la pomme de terre, malgré son nom. " +
		"Elle est dotée d'une très grande richesse en pro-vitamine A (ou bêta-carotène). " +
		"Elle fournit aussi une énergie de longue durée, grâce à l’amidon qu’elle renferme."
	trieString(s)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		var x int = rand.Intn(199) + 1
		fmt.Printf("%d - %d\n", x, syracuse(x))
	}

	num3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	tel := new(string)
	*tel = "450a1232cdg5672fw<112"
	telephone(tel)
	fmt.Println(*tel)
}

func trieString(s string) {
	s = strings.ToLower(s)
	r := strings.NewReplacer("'", " ", "(", "", ")", "", ".", "", ",", "")
	//re, _ := regexp.Compile(`[\.,\(\)]`)

	s = r.Replace(s)

	var words = strings.Fields(s)
	var element = make(map[string]int)
	for _, mot := range words {
		element[mot]++
	}
	fmt.Print(element)
}

/*
La conjecture de Syracuse: si le nombre est pair, on le divise par deux. Si le nombre est impair,

on le multiplie par 3 et on ajoute 1.

Par exemple, si on essaie avec 3, la suite sera: 3, 10, 5, 16, 8, 4, 2, 1. Donc nous avons 8 entiers.
*/
func syracuse(x int) int {
	if compte, ok := cache[x]; ok {
		return compte
	}
	if x%2 == 0 {
		x /= 2
	} else {
		x = x*3 + 1
	}

	if x == 1 {
		return 1
	} else {
		var count = syracuse(x) + 1
		cache[x] = count
		return count
	}
}

func num3(args ...int) {
	for x, num := range args {
		if x != 0 {
			args[0] += num
		}
	}
	fmt.Println(args[0])
}

func telephone(tel *string) {
	reg, err := regexp.Compile("[^0-9]")
	if err != nil {

	}
	*tel = reg.ReplaceAllString(*tel, "")

	if len(*tel) < 10 {
		log.Fatal("lol")
	}
	var normal string
	normal = (*tel)[:3]
	normal += " "
	normal += (*tel)[3:6]
	normal += "-"
	normal += (*tel)[6:10]
	if len(*tel) > 10 {
		normal += " poste "
		normal += (*tel)[10:]
	}
	*tel = normal
}
