package Array_test

import (
	"fmt"
	"testing"
)

// 定义一个接口,声明三个方法
// Len() 元素个数
// Less() 比较元素
// Swap() 交换元素
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(data Sorter) Sorter {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len() - pass; i++{
			if data.Less(i + 1, i) {
				data.Swap(i, i + 1)
			}
		}
	}
	return data
}

func TestSorterArray(t *testing.T) {
	data := []int{12, 45, 566, -9312, 423, 46, 123}
	res := IntArray(data)
	fmt.Println(Sort(res))
}