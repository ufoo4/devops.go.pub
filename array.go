package main

import "fmt"

func main() {
	var list []int64
	var list2 []string

	l := len(list2) // 0
	c := cap(list2) // 0

	list = []int64{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]
	l = len(list)                 // 5
	c = cap(list)                 // 5

	list = make([]int64, 0, 5) // []
	l = len(list)              // 0
	c = cap(list)              // 5

	list = make([]int64, 5, 5)
	list2 = make([]string, 5, 5)
	l = len(list2)
	c = cap(list2)

	fmt.Println(list2, l, c)

}
