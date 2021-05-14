package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
}

type Circle struct{
	radius float64
}

type Rectangualar struct{
	width, height float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r Rectangualar) area() float64{
	return r.height * r.width
}

func getArea(s Shape) float64{
	return s.area()
}


func main(){
	circle := Circle{50}
	rectangle := Rectangualar{50, 60}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}