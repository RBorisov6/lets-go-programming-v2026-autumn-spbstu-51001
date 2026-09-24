package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number1, number2, operator string
	fmt.Scan(&number1, &operator, &number2)

	number1, err1 := strconv.ParseFloat(number1, 64)
	number2, err2 := strconv.ParseFloat(number1, 64)
	
	if err1 != nil {
		fmt.Println("Invalid first operand")
	} else if err2 != nil {
		fmt.Println("Invalid second operand")
	} else if operator != "+", && operator != "-" && operator != "*" && operator != "/"{
		fmt.Println("Invalid operation")
	} else if operator == "/" && number2 == 0 {
		fmt.Println("Division by zero")
	}

	switch operator {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
	}
}