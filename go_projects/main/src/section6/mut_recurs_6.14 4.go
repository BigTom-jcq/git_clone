package main

import "fmt"

func main() {
	fmt.Printf("%d is even: is%t\n", 16, even(16))
	fmt.Printf("%d is odd: is%t\n", 17, odd(17))
	fmt.Printf("%d is odd: is%t\n", 18, odd(18))
}

func even(i int) bool {
	if i == 0 {
		return true
	}
	return odd(RevSign(i) - 1)
}

func odd(i int) bool {
	if i == 0 {
		return false
	}
	return even(RevSign(i) - 1)
}

func RevSign(i int) int {
	if i < 0 {
		return -i
	}
	return i
}