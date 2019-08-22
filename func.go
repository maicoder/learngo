package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(div(13, 3))
}
