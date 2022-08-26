package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	v := []int{133, 22, 31, 234, 45, 3, 23, 34, 12}
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	multiplyBy2()
	bubbleSort(v)
	bubbleSortB(v)
	min, max, moy := minMaxMoy(x)
	fmt.Printf("Min: %d, Max: %d, Moy: %f\n", min, max, moy)
	syracuse(3)
	fibonaci2()

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000)
	if isPrime(random) {
		fmt.Println("le nombre ", random, " est premier")
	} else {
		fmt.Println("le nombre ", random, " n'est pas premier")
	}

}

/*
## Numéro 1

	Demandez à l'utilisateur d'entrer un nombre réel, doublez le et affichez le résultat.

	Pas besoin de faire de gestion d'erreurs.
*/
func multiplyBy2() {
	var nb float32
	fmt.Println("Entrer un nombre decimal")
	fmt.Scanf("%f", &nb)
	fmt.Printf("%f\n", nb*2)
}

/*
## Numéro 2

Soit l'array suivant:
```go
v := [9]int{133, 22, 31, 234, 45, 3, 23, 34, 12}
```

Créer le code pour le trier avec un tri par bulle.

Utilisez des boucles for "standards" (avec les 3 parties)
*/
func bubbleSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			if v[i] > v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
	fmt.Printf("%v\n", v)
}

/*
## Numéro 2b

Refaites le numéro 2, mais utilisez des slices au lieu de faire des boucles sur le tableau original.
Utilisez aussi des boucles for avec un range.
*/
func bubbleSortB(v []int) {
	/*
		ma facon batarde

		for _ = range v {
			for j := range v {
				if j != len(v)-1 {
					s := v[j : j+2]
					if s[0] > s[1] {
						s[0], s[1] = s[1], s[0]
					}
				}
			}
		}
	*/

	v1 := v[:len(v)-1]
	for i := range v1 {
		v2 := v[i+1:]
		for j := range v2 {
			if v1[i] > v2[j] {
				v1[i], v2[j] = v2[j], v1[i]
			}
		}
	}
	fmt.Println(v)
}

/*
## Numéro 3

Soit l'array suivant:
```go
x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97, 9,17}
```

Faîtes une méthode minmaxmoy qui retourne le nombre plus petit et le plus grand et la moyenne des éléments qui ont été passés.
Appelez cette méthode là de votre fonction main.

	NB: Vous ne devez pas utiliser trois méthodes, mais une seule.
	L'appel de la fonction doit se faire ainsi:

```go
min, max, moy := minmaxmoy(x)
```
*/
func minMaxMoy(x []int) (int, int, float64) {
	min := x[0]
	max := x[0]
	moy := 0.0
	for _, v := range x {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		moy += float64(v)
	}
	moy /= float64(len(x))
	return min, max, moy
}

/*
## Numéro 4

Créez une fonction qui reçoit un entier en paramètre et qui calcule le nombre d'entiers dans la conjecture de Syracuse,

	du nombre reçu jusqu'au nombre 1. Choisissez un nombre (hardcodé) pour faire vos tests.

La conjecture de Syracuse: si le nombre est pair, on le divise par deux. Si le nombre est impair,

	on le multiplie par 3 et on ajoute 1.

Par exemple, si on essaie avec 3, la suite sera: 3, 10, 5, 16, 8, 4, 2, 1. Donc nous avons 8 entiers.
*/
func syracuse(n int) {
	x := []int{n}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		x = append(x, n)
	}
	fmt.Println(x)
}

/*
## Numéro 5

Affichez tous les nombres de la suite de Fibonacci pour les X premiers entiers (X est reçu en paramètre).
Choisissez un nombre (hardcodé) pour faire vos tests.

La suite de Fibonacci: Le nombre 0 est 0, le nombre 1 est 1. Pour chaque nombre suivant,
on additionne les deux nombres précédents. Pour le 7e nombre nous aurons: 1, 1, 2, 3, 5, 8, 13

Faire ce numéro de deux manières:

1. Faire un print pour chaque nombre trouvé
2. Mettre tous les nombres dans un tableau afin de faire les prints seulement lorsque tous les nombres sont trouvés
*/
func fibonaci1() {
	var n int
	fmt.Println("Entrer un nombre")
	fmt.Scan(&n)

	for i := 2; i < n; i++ {
		x = append(x, x[i-1]+x[i-2])
	}
	for _, v := range x {
		fmt.Println(v)
	}
}
func fibonaci2() {
	var n int
	fmt.Println("Entrer un nombre")
	fmt.Scan(&n)
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	for i := 2; i < n; i++ {
		x = append(x, x[i-1]+x[i-2])
	}
	for _, v := range x {
		fmt.Println(v)
	}
}

/*
	   ## Numéro 6

	   Inscrivez le code nécessaire pour créer un générateur de nombre premier
	   (vous générez aléatoirement le nombre et valider s'il est premier).

	   Pour générer un nombre entier aléatoire:

	   1. Importez les paquets "math/rand" et "time"
	   2. Initialisez le générateur de nombre aléatoire avec le code suivant (lisez bien le code pour le comprendre):
		`rand.Seed(time.Now().UnixNano())`
	   3. Générer un nombre aléatoire entre 0 et 1000 avec le code suivant: `rand.Intn(1000)`

	   Notez que la fonction "rand.Intn" ne prend que la borne max comme paramètre. Pour avoir un min il faut faire des maths ;)

	   La fonction "rand.Int" (sans le 'n') vous donne un entier positif (ou zéro), parmis tous les entiers possibles.
*/
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
