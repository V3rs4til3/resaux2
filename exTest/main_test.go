package exTest

import "testing"

type paires struct {
	param, reponse int
}

var valeursTest = []paires{
	{0, 0},
	{10, 55},
	{4, 3},
}

func TestFibonacciRecursion(t *testing.T) {
	for _, paire := range valeursTest {
		f := FibonacciRecursion(paire.param)
		if f != paire.reponse {
			t.Errorf("FibonacciRecursion(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}

func TestFibonacciLoop(t *testing.T) {
	for _, paire := range valeursTest {
		f := FibonacciLoop(paire.param)
		if f != paire.reponse {
			t.Errorf("FibonacciLoop(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}

func BenchmarkFibonacciLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciLoop(10)
	}
}

func BenchmarkFibonacciRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursion(10)
	}
}
