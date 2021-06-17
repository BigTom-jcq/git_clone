package goroutine_test

import (
	"fmt"
	"testing"
)

func printer(c chan int) {
	// 开始无限循环
	for {
		// 从channel中获取一个数据
		data := <-c
		// 将0视为结束
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	// 通知main已经循环结束(我搞定啦!) 第二步
	c <- 0
}

func TestConcurrencePrint(t *testing.T) {
	// 创建channel
	c := make(chan int)
	// 并发执行
	go printer(c)

	for i := 1; i <= 3; i++ {
		// 将数据通过channel投送给printer
		c <- i
	}
	// 通知并发的printer结束(没数据啦!) 第一步
	c <- 0
	// 等待printer结束(搞定喊我!) 第三步
	<-c
}
