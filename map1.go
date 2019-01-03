package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	x["key1"] = 11
	x["key2"] = 12
	fmt.Println(x["key"])
	fmt.Println(x["key1"])
	fmt.Println(x["key2"])
}
