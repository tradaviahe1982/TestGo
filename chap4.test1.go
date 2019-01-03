package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	//
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, "đây là số chẵn")
		} else {
			fmt.Println(j, "đây là số lẻ")
		}
	}
}
