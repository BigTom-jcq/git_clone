package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestRangeChannel(t *testing.T) {
	// 构建通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		// 从3 循环到 0
		for i := 3; i >= 0; i-- {
			// 发送3 到 0 之间的值
			ch <- i
			// 每次发送完等待
			time.Sleep(time.Second)
		}
	}()
	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)
		// 遇到0 时 结束循环
		if data == 0 {
			break
		}
	}
}
