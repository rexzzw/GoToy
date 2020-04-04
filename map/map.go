package main

import (
	"fmt"
)

func main() {
	mapA := map[string]bool{
		"A": false,
	}

	getB, ok := mapA["B"]
	fmt.Println("getB:", getB, "ok:", ok)

	c := mapA["C"]
	fmt.Println("c:", c)
}
