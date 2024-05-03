package main

import (
	"fmt"
	"math/rand/v2"
)

func createArray(len int) []int {
	array := make([]int, len)
	for i := 0; i < len; i++ {
		array[i] = rand.IntN(100)
	}
	return array
}
func hasDuplicate(nums []int) bool {
	used := make(map[int]bool)
	for _, num := range nums {
		if used[num] {
			return true
		}
		used[num] = true

	}
	return false
}

func main() {
	var lenNums int
	fmt.Print("Введи длину массива: ")
	fmt.Scan(&lenNums)
	nums := createArray(lenNums)
	fmt.Println("Сгенерированный массив: ", nums)
	fmt.Println("В массиве встречается повторение? ", hasDuplicate(nums))
}
