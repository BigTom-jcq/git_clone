package promotion_debug

import (
	"fmt"
	"reflect"
	"testing"
)

func findchildstr(a string) []int {
	var re = map[int]string{}
	dp := []int{}
	dp = append(dp, -1)
	for i := 0; i < len(a); i++ {
		dp = append(dp, 0)
	}
	for i := 0; i < len(a); i++ {
		if i == 0 {
			re[0] = a[0:1]
		} else {
			re[i] = a[0 : i+1]
		}

	}
	for k, v := range re {
		for i := 0; i < len(v); i++ {
			if len(v) == 1 {
				dp[k+1] = 0
			} else {
				if reflect.DeepEqual(v[0:i], v[len(v)-i-1:len(v)-1]) && i < len(v)-i {
					dp[k] = len(v[0:i])
				}
			}
		}
	}
	return dp[0:len(a)]
}

func restring(a string, b string)(int,int) {
	v := findchildstr(b)
	p := a
	i := 0
	for i<=len(a)-len(b){
		p = a[i:]
		j:=0
		for j<len(b){
			if p[j] != b[j]{
				if v[j+1]==0{
					i++
				}else{
					i = v[j+1]+j
				}
				break
			}
			j++
		}
		if j==len(b){
			break
		}
	}
	return i,i+len(b)-1
}

func TestRestring(t *testing.T) {
	fmt.Println(restring("abcdefg","fg"))
}