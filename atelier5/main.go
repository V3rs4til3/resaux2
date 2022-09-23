package main

func main() {

}
func Somme(numbers []int) int {
	somme := 0
	for _, i := range numbers {
		somme += i
	}
	return somme
}
func Moyenne(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(Somme(numbers) / len(numbers))
}
func StandardDeviation(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	moy := Moyenne(numbers)
	somme := 0
	for _, i := range numbers {
		somme += (i - int(moy)) * (i - int(moy))
	}
	return float64(somme / len(numbers))
}
func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, y := range numbers {
		if max < y {
			max = y
		}
	}
	return max
}
func Min(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, y := range numbers {
		if min > y {
			min = y
		}
	}
	return min
}
func Median(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	if len(numbers)%2 == 0 {
		return numbers[len(numbers)/2]
	} else {
		return (numbers[len(numbers)/2+1] + numbers[len(numbers)/2-1]) / 2
	}
}

func BubbleSort(v []int) []int {
	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			if v[i] > v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
	return v
}
