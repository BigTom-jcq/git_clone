package main

import (
	"fmt"
)

func main() {
	var f = Adder1()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder1() func(i int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}