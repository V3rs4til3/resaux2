package main

import "testing"

func BenchmarkavecRoutine($b *testing.B) {
	for i := 0; i < b.N; i++ {
		avecRoutine()
	}
}

