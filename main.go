package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("hello world 2")

	name, age := "Artur", 32

	// %d и %s - якоря
	c := fmt.Sprintf("My name is %s. And I'm %d years old", name, age)
	fmt.Println(c)
}
