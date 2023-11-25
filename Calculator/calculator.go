package main

import "fmt" 
 
func main() {
	fmt.Println("Select an operator:")
	
	fmt.Println("\n\tAddition: +")
	fmt.Println("\n\tSubtraction: -")
	fmt.Println("\n\tMultiplication: *")
	fmt.Println("\n\tDivision: /")

	fmt.Print("\nEnter an operation: ")
	var operator string
	fmt.Scanln(&operator)

	fmt.Print("\nEnter a number: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Print("\nEnter a second number: ")
	var num2 int
	fmt.Scanln(&num2)

	switch (operator) {
	case "+": add(num1,num2)
	case "-": subtract(num1,num2)
	case "*": multiply(num1,num2)
	case "/": divide(num1,num2)
	default: fmt.Println("Invalid operator entered.")
	}
}

func add(a int, b int) {
	var res int = a + b
    fmt.Println("\nAnswer is: ", res)
}

func subtract(a int, b int) {
	var res int = a - b
    fmt.Println("\nAnswer is: ", res)
}

func multiply(a int, b int) {
	var res int = a * b
    fmt.Println("\nAnswer is: ", res)
}

func divide(a int, b int) {
	var res int = a / b
    fmt.Println("\nAnswer is:", res)
}