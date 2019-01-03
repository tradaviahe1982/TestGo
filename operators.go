package main

import "fmt"

func main() {
	var m, n int = 100, 25
	var a, b string = "John", "Carry"
	var sum, difference, product, quotient, remainder int
	var sumstring string
	sum = m + n // +   sum   integers, floats, complex values, strings
	sumstring = a + " " + b
	difference = m - n // -   difference   integers, floats, complex values
	product = m * n    // *   product   integers, floats, complex values
	quotient = m / n   // /   quotient   integers, floats, complex values
	remainder = m % n  //%   remainder  integers
	fmt.Println(sum)
	fmt.Println(sumstring)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
	fmt.Println(remainder)
}
