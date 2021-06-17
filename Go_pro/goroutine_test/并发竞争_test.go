package goroutine_test

import (
  "fmt"
  "runtime"
  "testing"
)

var (
	count int32
	//wg    sync.WaitGroup
)

func inCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

func TestCompetition(t *testing.T) {
	wg.Add(2)
	go inCount()
	go inCount()
	wg.Wait()
	fmt.Println(count)
}
