package struct_study
// 方法和函数统一调用

import (
	"fmt"
	"testing"
)

// 声明一个结构体
type class struct {
}

// 给结构体添加Do方法
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

// 普通函数Do
func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func TestCall(t *testing.T) {
	// 声明一个函数回调
	var delegate func(int)

	// 创建结构体实例
	c := new(class)

	// 将回调设为c的Do方法
	delegate = c.Do

	// 调用
	delegate(100)

	// 将回调设为普通函数
	delegate = funcDo

	// 调用
	delegate(100)
}
