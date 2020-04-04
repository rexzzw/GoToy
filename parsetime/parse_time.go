package main

import (
	"fmt"
	"time"
)

func main() {
	timeString := "2019-01-01T00:01:02Z08:00"
	timeObj, err := time.Parse(timeString, time.RFC3339)
	if err != nil {
		fmt.Printf("err:%v", err)
	} else {
		fmt.Printf(timeObj.String())
	}
}
