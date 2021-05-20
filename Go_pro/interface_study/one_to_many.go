package main

import (
	"fmt"
	"io"
)

// 一个类型可以实现多个接口
type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("data is write")
	return 0, nil
}

func (s *Socket) Close() error {
	fmt.Println("data is close")
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)

}

type Closer interface {
	Close() error
}

func usingWriter(writer io.Writer)  {
	writer.Write(nil)
}

func usingCloser(closer io.Closer)  {
	closer.Close()
}

func main() {
	// 实例化Socket
	s := new(Socket)

	usingWriter(s)

	usingCloser(s)
}