package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.

type Writer interface {
	Write(data string) string
}

type ConsoleWriter struct {
}

func (c ConsoleWriter) Write(data string) string {
	fmt.Println("data >> ", data)
	return data
}

type ConsoleHashWriter struct {
}

func (w ConsoleHashWriter) Write(data string) string {
	fmt.Println(" data # ", data)
	return data
}

type Counter interface {
	Increment() int
	Decrement() int
}

type IntCounter int

func (c *IntCounter) Increment() int {
	*c = *c + 1
	return int(*c)
}

func (c *IntCounter) Decrement() int {
	*c = *c - 1
	return int(*c)
}

// Geometric interface example

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {

	return r.width * r.height

}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(c Geometry) {
	fmt.Println(c.area())
	fmt.Println(c.perim())
}

func main() {

	fmt.Println("Interfaces")

	var w Writer = ConsoleWriter{}
	w.Write("Hello")

	myInt := IntCounter(1)
	var inc Counter = &myInt
	fmt.Println(inc.Increment())
	fmt.Println(inc.Increment())
	fmt.Println(inc.Increment())
	fmt.Println(inc.Decrement())

	// Geomery example

	circle := Circle{radius: 5.6}
	rect := Rect{width: 400, height: 500}

	measure(circle)
	measure(rect)

	// Empty interfaces
	var myvar interface{} = "400"

	fmt.Println(myvar)

	switch myvar.(type) {
	case int, int32, int64:
		fmt.Println("Integer")
	case float64, float32:
		fmt.Println("Float type")
	case string:
		fmt.Println("String type")
	default:
		fmt.Println("Dont know")
	}

}
