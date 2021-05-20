package struct_study
// 非指针类型接收器
import (
	"fmt"
	"testing"
)

// 定义点结构
type Point struct {
	x int
	y int
}

// 非指针接收器的add方法
func (p Point) Add(other Point) Point {
	// 成员值与参数相加后返回新的结构
	return Point{
		p.x + other.x,
		p.y + other.y,
	}
}

func TestReceiverNotPointer(t *testing.T) {
	// 初始化点
	p1 := Point{1, 1}
	p2 := Point{10, 9}

	result := p1.Add(p2)
	fmt.Println(result)
}