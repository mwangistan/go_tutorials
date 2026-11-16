package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter 2 numbers: ")
	var a, b int
	fmt.Scan(&a, &b)
	quotient, remainder, err := divide(a, b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(a, "/", b, "=", quotient, "remainder =", remainder)
	}
}

func divide(a, b int) (int, int, error) {
	var err error
	var remainder int = 0
	var quotient int = 0
	if b == 0 {
		err = errors.New("division by zero")
		return quotient, remainder, err
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder, err
}
