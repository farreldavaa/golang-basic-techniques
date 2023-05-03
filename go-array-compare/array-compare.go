package main

import "fmt"

func main() {
	fmt.Println(compareArray([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(compareArray([]int{1, 2, 3}, []int{1, 2, 4}))
}

// using Ellipsis
func compareArray(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}
