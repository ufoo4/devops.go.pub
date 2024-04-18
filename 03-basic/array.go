package main

import "fmt"

func main() {
	var list []int64
	var list2 []string

	l := len(list) // 0
	c := cap(list) // 0
	l2 := len(list2)
	c2 := cap(list2)

	_ = l
	_ = c
	_ = l2
	_ = c2

	list = []int64{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]
	l = len(list)                 // 5
	c = cap(list)                 // 5

	list = make([]int64, 0, 5) // []
	l = len(list)              // 0
	c = cap(list)              // 5

	list = make([]int64, 5, 5)   // [0,0,0,0,0]
	list2 = make([]string, 5, 5) // ["","","","",""]

	l = len(list)
	c = cap(list)

	list = []int64{1, 2, 3, 4} // [1,2,3,4]
	list = append(list, 5)     // [1,2,3,4,5]

	list = make([]int64, 0, 3)
	list = append(list, 1) // [1] len:1 cap:3
	list = append(list, 2) // [1,2] len:2 cap:3
	list = append(list, 3) // [1,2,3] len:3 cap:3
	list = append(list, 4) // [1,2,3,4] len:4 cap:6
	list = append(list, 5) // [1,2,3,4,5] len:5 cap:6

	l = len(list)
	c = cap(list)

	fmt.Println(list, l, c)

}
