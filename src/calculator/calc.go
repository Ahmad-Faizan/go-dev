package main

import (
	"fmt"
	"math"
	"os"
)

func add(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return math.Abs(x - y)
}

func mul(x, y float64) float64 {
	return x * y
}

func div(x, y float64) float64 {
	return x / y
}

func main() {
	var a, b float64
	var opr string

	fmt.Println("Enter the operation (+ , - , * , / ) ")
	fmt.Scan(&opr)

	fmt.Println("Enter the two operands ")

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Invalid number format", err)
		os.Exit(1)
	}
	switch opr {
	case "+":
		fmt.Printf("Result of %v operation : %.3f", opr, add(a, b))
	case "-":
		fmt.Printf("Result of %v operation : %.3f", opr, sub(a, b))
	case "*":
		fmt.Printf("Result of %v operation : %.3f", opr, mul(a, b))
	case "/":
		fmt.Printf("Result of %v operation : %.3f", opr, div(a, b))
	default:
		fmt.Println("Invalid Operator specified")
	}
}
