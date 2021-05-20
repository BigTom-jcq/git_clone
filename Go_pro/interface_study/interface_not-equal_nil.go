package main

import "fmt"

// 定义一个结构体
type MyImplement struct {
}

// 实现fmt.Stringer 的 String方法
func (m *MyImplement) String() string {
	return "hi"
}

// 在函数中返回fmt.Stringer接口
func GetStringer() fmt.Stringer {
	// 赋值nil
	var s *MyImplement = nil

	// 发现nil类型值返回时直接返回nil
	if s == nil {
		return nil
	}

	return s
}

func main() {
	// 判断返回值是否为nil
	if GetStringer() == nil {
		fmt.Println("GetStringer == nil")
	} else {
		fmt.Println("GetStringer != nil")
	}
}