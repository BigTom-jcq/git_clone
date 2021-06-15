package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}

//func main() {
//	a := [...]string{"a", "b", "c", "d"}
//	for i := range a {
//		fmt.Println("Array item", i, "is", a[i])
//	}
//}
//func main() {
//	var arr1 = new([5]int)
//	arr2 := *arr1
//	arr2[2] = 100
//	fmt.Printf("arr1 is %v\n", arr1)
//	fmt.Printf("arr2 is %v\n", arr2)
//}