package main

import "fmt"

func increment(p *int) {
	*p++
}

func main() {
	num := 10
	increment(&num)
	fmt.Println("After increment:", num)
}
