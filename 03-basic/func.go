package main

import "fmt"

func addPredix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}
func addPrefixWithLen(origin string) (res string, lenght int) {
	res = "prefix_" + origin
	lenght = len(res)
	return res, lenght
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return factorial(n-1) * n
}
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	myString := addPredix("my_string")
	fmt.Println(myString) // prefix_my_string

	myString, err := addPrefixWithErr("error_string")
	fmt.Println(err)      // <nil>
	fmt.Println(myString) //prefix_error_string

	myFactorial := factorial(65)
	fmt.Println(myFactorial)

	myAdder := adder()
	fmt.Println(myAdder)
}
