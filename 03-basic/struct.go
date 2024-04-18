package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
		// Z string
		// T[]int64
	}
	p := Point{
		X: 5,
		Y: 19,
	}
	p = Point{7, 23}

	fmt.Println(p.X) // 7
	fmt.Println(p.Y) // 23

	p.X = 12332

	fmt.Println(p.X) // 12332

}
