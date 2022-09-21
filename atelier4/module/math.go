package math

func Somme(numbers []int) int {
	somme := 0
	for _, i := range numbers {
		somme += i
	}
	return somme
}
func Moyenne(numbers []int) float64 {
	return float64(Somme(numbers) / len(numbers))
}
func StandardDeviation(numbers []int) float64 {
	moyenne := Moyenne(numbers)
	var somme float64
	for _, i := range numbers {
		somme += float64(i - (int(moyenne))*i - (int(moyenne)))
	}
	return somme / float64(len(numbers))
}
func Max(numbers []int) int {
	max := numbers[0]
	for _, y := range numbers {
		if max < y {
			max = y
		}
	}
	return max
}
func Min(numbers []int) int {
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
