package main

import (
	"fmt"
	"time"
)

const LMT = 41

var fibs [LMT]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LMT; i++ {
		result = fibonacci1(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("fibonacci() time is: %s\n", delta)
}

func fibonacci1(i int) (res uint64) {
	if fibs[i] != 0 {
		res = fibs[i]
		return
	}
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci1(i-1) + fibonacci1(i-2)
	}
	fibs[i] = res
	return
}
