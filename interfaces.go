package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area1() float64
	perim1() float64
}

type rect1 struct {
	width, height float64
}
type circle1 struct {
	radius float64
}

func (r rect1) area1() float64 {
	return r.height * r.width
}

func (r rect1) perim1() float64 {
	return 2*r.width + 2*r.height
}

func (c circle1) area1() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle1) perim1() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area1())
	fmt.Println(g.perim1())
}

func interfaces() {
	r := rect1{3, 4}
	c := circle1{5}

	// measure(&r) still works fine, but not measure(c)
	// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	measure(r)
	measure(&c)
}
