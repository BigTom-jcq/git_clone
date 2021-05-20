package Array_test

import (
	"fmt"
	"testing"
)

// 最大盛水容器
// 双指针解法

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height) - 1
	for i != j {
		hi, hj := height[i], height[j]
		s := min(hi, hj) * (j - i)
		if s > res {
			res = s
		}
		if hi > hj {
			j--
		} else {
			i++
		}
 	}
	return res
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}