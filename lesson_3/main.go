package main

import "fmt"

type Point struct {
	A string
	B int
}

func (p *Point) method() {
	fmt.Println(p.A)
}

func main() {
	pointer()

	p1 := Point{
		A: "hello",
		B: 1,
	}

	fmt.Println(p1)
	p1.method()
	p2 := &p1
	fmt.Println(p2)
	p2.method()
}

func pointer() {
	a := "hello world"

	p := *&a
	fmt.Println(p)
}
