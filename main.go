package main

import "fmt"

type rectangle struct {
	width, height float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	r := rectangle{23.3, 90}
	fmt.Print(r.area())
}
