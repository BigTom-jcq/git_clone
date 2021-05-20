package promotion_debug

import (
	"fmt"
	//"Go_pro/promotion/promotion_debug"
	"testing"
)

func TestRule(t *testing.T)  {
	info,err:=Rule()
	if err!=nil{
		fmt.Println("解码失败", err.Error())
	}else {
		fmt.Println(info)
	}

}
