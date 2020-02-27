package main

import (
    "fmt"
)

func main() {
    status := "0"
    switch status {
    case "0", "abc":
        fmt.Println("case 0")
    }
}
