package	struct_study

import (
	"fmt"
)

type Command struct {
	Name string
	Var *int
	Comment string
}

var version int = 1

func StructTest1() {
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
	fmt.Println(cmd.Name)
	fmt.Println(cmd.Var)
	fmt.Println(cmd.Comment)
}