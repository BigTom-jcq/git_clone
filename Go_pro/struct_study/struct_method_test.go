package struct_study
// 结构体方法
import (
	"fmt"
	"testing"
)

type Bag struct {
	items []int
}

func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}

func TestStruct(t *testing.T) {
	//b := &Bag{}
	b := new(Bag)
	b.Insert(1001)
	fmt.Println(b.items)
}