package main

import (
	"fmt"

	"github.com/PrashantSabnekar/gotestpackage"
)

func main() {

	fmt.Println("Using GoDemoPackage")

	var x, y, res int
	x = 100
	y = 200
	res = gotestpackage.Add(x, y)
	fmt.Println("Result = ", res)

}
