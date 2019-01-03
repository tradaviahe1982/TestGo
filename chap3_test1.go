package main

import "fmt"

func main() {
	var (
		a = 5
		b = 6
		c = 7
	)
	const y string = "Hello World!"
	var x string = "Xin chào bạn Nguyễn Việt Anh!"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a, b, c)
	fmt.Print("Nhập một số: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
