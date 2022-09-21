package main

import (
	"fmt"
	"math"
)

type Cercle struct {
	x, y, r float64
}

func (c *Cercle) aire() float64 {
	return math.Pi * c.r * c.r
}

func (c *Cercle) perimetre() float64 {
	return math.Pi * 2 * c.r
}

type Rectangle struct {
	x, y float64
}

func (r *Rectangle) aire() float64 {
	return r.x * r.y
}

func (r *Rectangle) perimetre() float64 {
	return (r.x + r.y) * 2
}

type Carre struct {
	h float64
}

func (r *Carre) aire() float64 {
	return r.h * r.h
}

func (r *Carre) perimetre() float64 {
	return r.h * 4
}

type Forme interface {
	aire() float64
	perimetre() float64
}

func aireEtPerimetreTotal(formes ...Forme) (float64, float64) {
	var totalA, totalP float64

	for _, i := range formes {
		totalA += i.aire()
		totalP += i.perimetre()
	}

	return totalA, totalP
}

func main() {
	c := Cercle{0, 0, 5}
	c1 := Cercle{2, 4, 12}
	r := Rectangle{4, 5}
	r1 := Rectangle{8, 12}
	cr := Carre{4}
	cr1 := Carre{6}
	aire, perimetre := aireEtPerimetreTotal(&c, &c1, &r, &r1, &cr, &cr1)
	fmt.Println("aire : ", aire, " et perimetre : ", perimetre)
}
