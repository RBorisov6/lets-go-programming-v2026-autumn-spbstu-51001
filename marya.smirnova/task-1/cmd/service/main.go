package main

import (
	"fmt"
)

func main() {
	var (
		x1, x2 int
		op     string
	)
	_, err1 := fmt.Scan(&x1)
	_, err2 := fmt.Scan(&x2)
	_, err3 := fmt.Scan(&op)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}
	switch op {
	case "+":
		fmt.Println(x1 + x2)
	case "-":
		fmt.Println(x1 - x2)
	case "*":
		fmt.Println(x1 * x2)
	case "/":
		if x2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(x1 / x2)
	default:
		fmt.Println("Invalid operation")
	}
}
