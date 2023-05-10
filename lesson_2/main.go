package main

import "fmt"

const TG_IP string = "aolo"

func main() {
	// s, s2 := test()
	// s3 := test1()
	// fmt.Println(s, s2, s3)

	sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += 1
	// }
	for sum < 10 { // while
		sum += 1
	}
	fmt.Println(sum)
}

func test() (a, b string) {
	a = "hello"
	b = "world"

	return a, b
}

func test1() string {
	return "empty"
}
