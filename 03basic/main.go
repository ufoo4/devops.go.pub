package main

import "fmt"

func main() {
	var x1 [5]int
	var x2 [0]int
	var x3 [0]string

	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{4, 5, 6}
	arr3 := [3]int{7, 8, 9}

	//arr3 = [2]int{1, 2}
	arr5 := [...]int32{1, 2, 3, 4, 5, 6}
	s := len(arr5)
	fmt.Println(x1, x2, x3, arr, arr2, arr3, arr5, s)
}
