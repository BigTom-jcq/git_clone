package main

import "fmt"

func main() {
	//var arrAge = [5]int{3, 4, 6, 89, 123}
	//var arrLazy = [...]int{4, 5, 6, 7, 21}
	//var arrLazy = []int{4, 5, 6, 7, 21}
	var arrKeyValue = [5]string{3: "Bob", 4: "Tom"}
	//var arrKeyValue = []string{3: "Bob", 4: "Tom"}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
