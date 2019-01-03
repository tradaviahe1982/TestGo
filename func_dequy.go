package main

import "fmt"

func hamdequy(x int) int {
	if x == 0 {
		return 1
	}
	return x * hamdequy(x-1)
}
func main() {
	fmt.Println(hamdequy(5))
}
