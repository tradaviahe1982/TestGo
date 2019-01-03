package main

func f() (int, int, float64) {
	return 5, 6, 89.99
}
func main() {
	x, y, z := f()
	println(x, y, z)
}
