package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) int { return a + b }
func Multiply(a, b int) int { return a * b }

func Divide(a, b int) (int, error) {
	if b == 0 { return 0, errors.New("division by zero") }
	return a / b, nil
}

func main() {
	fmt.Println("Add(2, 3) =", Add(2, 3))
	Divide(10, 0)
}
