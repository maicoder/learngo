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

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)
}
