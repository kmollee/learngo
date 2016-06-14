/*
example of variable scope
*/
package main

import "fmt"
import "math"

type Retangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

/*
extend struct function
http://golangtutorials.blogspot.tw/2011/06/methods-on-structs.html
*/
func (rect Retangle) area() float64 {
	return rect.width * rect.height
}

// normal function
func retangleArea(rect Retangle) float64 {
	return rect.width * rect.height
}

type Circle struct {
	radius float64
}

func (cir Circle) area() float64 {
	return math.Pow(cir.radius, 2) * math.Pi
}

// the interface of area
type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	num := 3
	doubleNum := func() int {
		num *= 2
		return num
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	// use parameter assign
	rect1 := Retangle{leftX: 0, topY: 50, height: 10, width: 10}
	// use order to assign
	rect2 := Retangle{0, 0, 50, 50}
	fmt.Println("Retangle's area is", retangleArea(rect1))
	fmt.Println("Retangle2's area is", rect2.area())

	rt := Retangle{0, 0, 20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle Area=", getArea(rt))
	fmt.Println("Circle Area=", getArea(circ))
}
