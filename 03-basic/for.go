package main

import "fmt"

func main() {
	sl := []int64{9, 8, 7, 1}
	for key, value := range sl {
		fmt.Printf("key: %v, val: %v \n", key, value)
	}
	// key: 0, val: 9
	// key: 1, val: 8
	// key: 2, val: 7
	// key: 3, val: 1

	for _, value := range sl {
		fmt.Printf("val: %v \n", value)
	}
	// val: 9
	// val: 8
	// val: 7
	// val: 1
	ages := map[string]int{
		"Андрей":    30,
		"Анастасия": 25,
	}
	for key, value := range ages {
		fmt.Printf("Имя: %v; Лет: %v \n", key, value)
	}
	// Имя: Андрей; Лет: 30
	// Имя: Анастасия; Лет: 25

	word := "слёрм"
	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
		fmt.Printf("%T", word[i])
	}
	// странно, почему так?
	// 209
	// uint8129
	// uint8208
	// uint8187
	// uint8209
	// uint8145
	// uint8209
	// uint8128
	// uint8208
	// uint8188
	// uint8%

	for key, value := range word {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Printf("%T", value)
	}
	//1089
	//int322
	//1083
	//int324
	//1105
	//int326
	//1088
	//int328
	//1084
	//int32%
}
