package main

import (
	"fmt"
	"math"
)

type shaper interface {
	findArea() float64
	findCircumference() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) findArea() float64 {
	return s.length * s.length
}

func (s square) findCircumference() float64 {
	return s.length * 4
}

func (c circle) findArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) findCircumference() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shaper) {
	fmt.Printf("Area of %T is: %0.2f\n", s, s.findArea())
	fmt.Printf("Circumference of %T is: %0.2f\n", s, s.findCircumference())
}

func main() {
	shapes := []shaper{
		square{length: 6.33},
		circle{radius: 5.66},
		circle{radius: 14.1},
		square{length: 24.1},
	}

	for _, value := range shapes {
		printShapeInfo(value)
		fmt.Println("------")
	}
}