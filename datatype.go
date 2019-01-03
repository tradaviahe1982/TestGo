package main

import "fmt"

//Global variable declaration
var (
	m int
	n int
)

func main() {
	var x int = 1 // Integer Data Type
	var y int     //  Integer Data Type
	fmt.Println(x)
	fmt.Println(y)

	var a, b, c = 5.25, 25.25, 14.15 // Multiple float32 variable declaration
	fmt.Println(a, b, c)

	city := "Berlin"     // String variable declaration
	Country := "Germany" // Variable names are case sensitive
	fmt.Println(city)
	fmt.Println(Country) // Variable names are case sensitive

	food, drink, price := "Pizza", "Pepsi", 125 // Multiple type of variable declaration in same line
	fmt.Println(food, drink, price)
	m, n = 1, 2
	fmt.Println(m, n)
}
