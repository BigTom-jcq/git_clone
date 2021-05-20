package promotion_debug

import (
	"encoding/json"
	"fmt"
	"os"
)

type rule struct {

}
func Rule()  ([]rule,error){
	var info [] rule
	filePtr, err := os.Open("/Users/jiang/Desktop/code/Go_pro/promotion/promotion_debug/rule.json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s\n]", err.Error())
		return info, err
	}
	defer filePtr.Close()

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
	return info,err
}
