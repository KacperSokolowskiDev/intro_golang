package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width float64
	height float64
}

func (c *circle) area() float64 {
	return c.radius * 2 * math.Pi
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

func someFunction(someShape shape) {
	fmt.Println("Area of the shape:")
	fmt.Println(someShape.area())
}

func main() {

	circle1 := circle{radius: 23.45}
	rect1 := rectangle{width: 87.2, height: 22.435}

	someFunction(&circle1) //must send the pointer to the shape
	someFunction(&rect1)

}