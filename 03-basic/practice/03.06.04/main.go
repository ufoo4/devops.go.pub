package main

import "fmt"

func countCharacters(str string) map[rune]int {
	mp := make(map[rune]int)
	for _, v := range str {
		mp[v]++
	}
	return mp
}
func main() {
	var str string = "съешь ещё этих мягких французских булок, да выпей чаю"

	for i, v := range countCharacters(str) {
		fmt.Printf("%q - %d\n", i, v)
	}

}
