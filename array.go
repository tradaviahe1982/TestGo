package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	//
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83
	//
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	//
	fmt.Println("Trung bình: ", total/5)
	//
	var total1 float64 = 0
	for _, value := range y {
		total1 += value
	}
	fmt.Println("Trung bình là: ", total1/5)
}
