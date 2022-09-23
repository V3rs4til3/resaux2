package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type paires = struct {
	param   []int
	reponse int
}

var valeurTestSum = []paires{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 28},
	{[]int{5, 6, 7, 8, 9, 10, 11}, 56},
	{[]int{11, 12, 13, 14, 15, 16}, 81},
	{[]int{}, 0},
	{[]int{0}, 0},
}

func TestSomme(t *testing.T) {
	for _, paire := range valeurTestSum {
		f := Somme(paire.param)
		if f != paire.reponse {
			t.Errorf("Somme(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}

type paireM = struct {
	param   []int
	reponse float64
}

var valeurTestMoy = []paireM{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
	{[]int{5, 6, 7, 8, 9, 10, 11}, 8},
	{[]int{11, 12, 13, 14, 15, 16}, 13.5},
	{[]int{}, 0},
	{[]int{0}, 0},
}

func TestMoyenne(t *testing.T) {
	for _, paire := range valeurTestMoy {
		f := Moyenne(paire.param)
		if f != float64(paire.reponse) {
			t.Errorf("Moyenne(%d) = %f; want %f", paire.param, f, paire.reponse)
		}
	}
}

var valeurTestDev = []paireM{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 2},
	{[]int{5, 6, 7, 8, 9, 10, 11}, 2},
	{[]int{11, 12, 13, 14, 15, 16}, 2.5},
	{[]int{}, 0},
	{[]int{0}, 0},
}

func TestStandardDeviation(t *testing.T) {
	for _, paire := range valeurTestDev {
		f := StandardDeviation(paire.param)
		if f != float64(paire.reponse) {
			t.Errorf("StandardDeviation(%d) = %f; want %f", paire.param, f, paire.reponse)
		}
	}
}

var valeurTestmax = []paires{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
	{[]int{5, 6, 7, 8, 9, 10, 11}, 11},
	{[]int{11, 12, 13, 14, 15, 16}, 16},
	{[]int{}, 0},
	{[]int{0}, 0},
}

func TestMax(t *testing.T) {
	for _, paire := range valeurTestmax {
		f := Max(paire.param)
		if f != paire.reponse {
			t.Errorf("Max(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}

var valeurTestmin = []paires{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 1},
	{[]int{5, 6, 7, 8, 9, 10, 11}, 5},
	{[]int{11, 12, 13, 14, 15, 16}, 11},
	{[]int{}, 0},
	{[]int{0}, 0},
}

func TestMin(t *testing.T) {
	for _, paire := range valeurTestmin {
		f := Min(paire.param)
		if f != paire.reponse {
			t.Errorf("Min(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}

var valeurTestMedian = []paires{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
	{[]int{5, 6, 7, 8, 9, 10, 11}, 8},
	{[]int{11, 12, 13, 14, 15, 16}, 13},
	{[]int{}, 0},
	{[]int{0}, 0},
}

func TestMedian(t *testing.T) {
	for _, paire := range valeurTestMedian {
		f := Median(paire.param)
		if f != paire.reponse {
			t.Errorf("Median(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}

func BenchmarkMoyenne(b *testing.B) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(1000000)
	for i := 0; i < b.N; i++ {
		Moyenne(permutation)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(permutation)
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(10000)
	sort.Ints(permutation)
}
