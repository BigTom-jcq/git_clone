package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i <= 25; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("fibonacci() time is: %s\n", delta)
}

func fibonacci(i int) (res int) {
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i - 1) + fibonacci(i - 2)
	}
	return
}