package main

import "fmt"

type gas struct {
	name  string
	price float64
}

type fuel interface {
	getGasInfo() string
}

func printGasInfo(f fuel) {
	fmt.Println(f.getGasInfo())
}

func (g gas) getGasInfo() string {
	return fmt.Sprintf("Name: %s, Price: %.2f", g.name, g.price)
}

func main() {
	myGas := gas{"Gasoline", 3.99}

	myGas.price = 4.29
	printGasInfo(myGas)
}
