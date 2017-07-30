package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
}

type Circle struct {
	Shape
	radius float64
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.radius
}

type Rectangle struct {
	Shape
	l, w float64
}

func (r *Rectangle) perimeter() float64 {
	return 2*r.l + 2*r.w
}

type Multishape struct {
	shapes []Shape
}

func main() {
	c := new(Circle)
	c.radius = float64(10)
	r := new(Rectangle)
	r.l = float64(8)
	r.w = float64(7)
	ms := new(Multishape)
	ms.shapes = append(ms.shapes, c, r)
	for _, s := range ms.shapes {
		fmt.Println(s.perimeter())
	}
}
