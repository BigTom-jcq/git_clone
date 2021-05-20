package struct_study
// 指针类型接收器
import (
	"fmt"
	"testing"
)

// 定义属性结构
type Property struct {
	value int  // 属性值
}

// 设置属性值
func (p *Property) SetValue(v int) {
	// 修改p的成员变量
	p.value = v
}

// 获取属性值
func (p *Property) Value() int {
	return p.value
}

func TestGetValue(t *testing.T) {
	p := new(Property)

	p.SetValue(100)

	fmt.Println(p.Value())
}