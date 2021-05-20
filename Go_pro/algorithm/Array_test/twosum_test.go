package Array_test

import (
	"fmt"
	"testing"
)

//class Solution:
//"""
//两数求和
//给定一个有序数组和一个目标值,当列表中的两个数相加等于目标值时,返回这两个数的下标
//"""
//	def twoSum(self, nums, target):
//		n = len(nums)
//		a = []
//		first = 0
//		last = n - 1
//		while first <= last:
//			if nums[first] + nums[last] == target:
//				a.append([first, last])
//			last -= 1
//			first += 1
//		return a

var num [][]int

func TwoSum(nums []int, target int) [][]int {
	last := len(nums) - 1
	for i := 0; i < last; i++ {
		if nums[i] + nums[last] == target {
			num = append(num, []int{i, last})
			last--
		}
	}
	return num
}


func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t1 := 9
	fmt.Println(TwoSum(nums, t1))
}
