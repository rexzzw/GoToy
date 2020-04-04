package main

import (
	"fmt"
	"reflect"
)

type Keys struct {
	Key1 string
	Key2 bool
	Key3 int64
}

func main() {
	keyObj := Keys{Key1: "string", Key2: true, Key3: 1}
	kk := reflect.ValueOf(keyObj)
	fmt.Println(kk)
}
