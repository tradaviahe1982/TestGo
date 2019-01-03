package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 0:
		fmt.Println("Số KHÔNG")
	case 1:
		fmt.Println("SỐ MỘT")
	case 2:
		fmt.Println("SỐ HAI")
	default:
		fmt.Println("KHÔNG RÕ SỐ NÀO")
	}
}
