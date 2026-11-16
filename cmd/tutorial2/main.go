package main

import "fmt"

func main() {
	var myNum int
	myNum = 10
	fmt.Println("myNum:", myNum)

	var myFloat float64
	myFloat = 3.14
	fmt.Println("myFloat:", myFloat)

	var myString string
	myString = "Hello"
	fmt.Println("myString:", myString)

	var myBool bool
	myBool = true
	fmt.Println("myBool:", myBool)

	var myChar rune
	myChar = 'A'
	fmt.Println("myChar:", myChar)

	var myByte byte
	myByte = 'B'
	fmt.Println("myByte:", myByte)

	var test = 42
	fmt.Println("test:", test)

	myNewString := "Hello, World!"
	fmt.Println("myNewString:", myNewString)

	const myConst = 100
	fmt.Println("myConst:", myConst)
}
