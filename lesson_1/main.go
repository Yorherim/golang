package main

import "fmt"

type Point struct {
	x string
	y string
}

func (p Point) log() {
	fmt.Println(p.x)
	fmt.Println(p.y)
}

func main() {
	p := &Point{
		x: "sasha",
		y: "pi",
	}
	fmt.Println(p)
	p1 := p
	p1.x = "dad"
	p.log()
}
