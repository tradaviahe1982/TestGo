package main

import "fmt"

func trungbinh(mang []float64) float64 {
	total := 0.0
	for _, v := range mang {
		total += v
	}
	return total / float64(len(mang))
}

func main() {
	mang := []float64{98, 93, 77, 82, 83}
	fmt.Println(trungbinh(mang))
}
