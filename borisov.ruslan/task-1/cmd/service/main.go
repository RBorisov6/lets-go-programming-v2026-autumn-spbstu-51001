package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1, input2, operator string
	fmt.Scan(&input1, &input2, &operator)

	number1, err1 := strconv.ParseFloat(input1, 64)
	number2, err2 := strconv.ParseFloat(input2, 64)
	
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	} else if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	} else if operator != "+" && operator != "-" && operator != "*" && operator != "/"{
		fmt.Println("Invalid operation")
		return
	} else if operator == "/" && number2 == 0 {
		fmt.Println("Division by zero")
		return
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