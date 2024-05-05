package main

import "fmt"

func sort(sls []string) bool {
	for i := 1; i < len(sls); i++ {
		if sls[i] < sls[i-1] {
			return false
		}
	}
	return true
}
func main() {
	sls := []string{"Close", "Danke", "Ferder", "Liben"}
	fmt.Println(sort(sls))
}
