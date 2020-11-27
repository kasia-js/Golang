package main

import (
	"fmt"
)

func Multiply(a, b int) int {
  return a * b
}

func main() {
	result := Multiply(2,3)
	fmt.Printf("%d", result)
}