package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]int{
		"x": 1,
		"y": 2,
	}
	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)

	fmt.Println(p1)
}
