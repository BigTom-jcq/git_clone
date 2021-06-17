package goroutine_test

import (
	"fmt"
	"testing"
)

func TestChannel1(t *testing.T) {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")

		// 通过通道通知test的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")
}
