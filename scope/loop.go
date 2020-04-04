package main

import (
	"fmt"
)

func f1(a int) (int, int) {
	fmt.Println("input", a)
	return a + 1, a + 2
}

func main() {
	var a = 0
	var b = 0
	for range []int{0, 1} {
		a, b = f1(a)
		fmt.Println(a, b)
	}
	fmt.Println(a)
	var open, close = 0, 1
	fmt.Println(open, close)
}
