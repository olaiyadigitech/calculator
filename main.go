package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	// Get the first number
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	// Get the operator
	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(&operator)

	// Get the second number
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Perform the calculation
	switch operator {
	case "+":
		fmt.Println("Result:", num1+num2)

	case "-":
		fmt.Println("Result:", num1-num2)

	case "*":
		fmt.Println("Result:", num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			fmt.Println("Result:", num1/num2)
		}

	default:
		fmt.Println("Invalid operator.")
	}
}
