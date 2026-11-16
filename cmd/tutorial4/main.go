package main

import "fmt"

func main() {
	var myArr [3]int = [3]int{1, 2, 3}
	fmt.Println(myArr)

	myarr2 := [3]int{4, 5, 6}
	fmt.Println(myarr2)

	mySlice := []int{7, 8, 9}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 10)
	fmt.Println(mySlice)

	myslice2 := make([]int, 3)
	fmt.Println(myslice2)

	var myMap map[string]int = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	_, ok := myMap["three"]
	fmt.Println(ok)

	for key, value := range myMap {
		fmt.Println(key, value)
	}

}
