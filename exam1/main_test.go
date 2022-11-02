package main

import "testing"

type paires struct {
	param, reponse int
}

var valeursTest = []paires{
	{0, 0},
	{50, 7},
	{64, 8},
	{4, 2},
	{-1, 0},
}

func TestJusquauCarre(t *testing.T) {
	for _, paire := range valeursTest {
		f := jusquauCarre(paire.param)
		if f != paire.reponse {
			t.Errorf("jusquauCarre(%d) = %d; want %d", paire.param, f, paire.reponse)
		}
	}
}
