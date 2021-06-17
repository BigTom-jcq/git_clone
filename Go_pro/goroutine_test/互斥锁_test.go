package goroutine_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var (
	mutex sync.Mutex
	wg sync.WaitGroup
	counter int64
)

func incCounter1(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock() // 释放锁,允许其他正在等待的goroutine进入临界区
	}
}

func TestMutexLock(t *testing.T) {
	wg.Add(2)

	go incCounter1(1)
	go incCounter1(2)

	wg.Wait()
	fmt.Println(counter)
	fmt.Println("=-=-=-=-=-=-=-=-=")
	fmt.Println(runtime.NumCPU())
}

// runtime.GOMAXPROCS(逻辑CPU数量)
// runtime.GOMAXPROCS(runtime.NumCPU())
