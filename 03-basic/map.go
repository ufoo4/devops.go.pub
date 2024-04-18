package main

import "fmt"

func main() {
	var m1 map[int32]bool
	var m2 map[string]string
	m3 := make(map[int]int)
	ages := map[string]int{
		"Андрей":    30,
		"Анастасия": 25,
		// ...
	}

	_ = m1
	_ = m2
	_ = m3

	age := ages["Андрей"] // 30
	_ = age
	ages["Наталья"] = 31        // map[Анастасия:25 Андрей:30 Наталья:31]
	fmt.Println(ages["Михаил"]) // 0

	fmt.Println(ages)
}
