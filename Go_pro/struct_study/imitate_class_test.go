package struct_study

import (
	"fmt"
	"testing"
)

type Flying struct {}

func (f Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct {}

func (w Walkable) Walk() {
	fmt.Println("can walk")
}

// 人类
type Human struct {
	Walkable
}

// 鸟
type Bird struct {
	Walkable
	Flying
}

func TestImitateClass(t *testing.T) {
	// 实例化鸟类
	bird := new(Bird)
	fmt.Println("Bird: ")
	bird.Fly()
	bird.Walk()

	// 实例化人类
	human := new(Human)
	fmt.Println("Human: ")
	human.Walk()
}