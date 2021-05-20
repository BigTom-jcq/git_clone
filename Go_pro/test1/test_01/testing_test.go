package test_01

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}
}

func TestAbs(t *testing.T) {
  t.Log("测试git")
}

func TestGit(t *testing.T) {
  t.Log("测试fetch")
}

func TestLocal(t *testing.T) {
  t.Log("local测试fetch")
}

func TestGit2(t *testing.T) {
}
