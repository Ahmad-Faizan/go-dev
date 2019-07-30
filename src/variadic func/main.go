package main

import "fmt"

func add(args ...float64) (sum float64) {
	for _, i := range args {
		sum += i
	}
	return
}
func main() {
	const pi = 3.14159
	const two = 2
	const five = 5
	const ten = 10.0

	fmt.Println(add(pi, two))
	fmt.Println(add(pi, two, five))
	fmt.Println(add(pi, two, five, ten))
	fmt.Println(add(2.3, 5.3, 8.3, 9.0, 92.1))
}
