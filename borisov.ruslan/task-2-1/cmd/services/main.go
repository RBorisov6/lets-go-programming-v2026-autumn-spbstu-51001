package main

import (
	"fmt"
)

func main() {
	const (
		mincount = 1
		maxcount = 1000
	)

	var N int

	_, err := fmt.Scan(&N)

	if err != nil || N < mincount || N > maxcount {
		fmt.Println("Invalid count of departments")
		return
	}

	for dep := 0; dep < N; dep++ {
		var K int

		_, err := fmt.Scan(&K)
		if err != nil || K < mincount || K > maxcount {
			fmt.Println("Invalid count of employers")
			return
		}

		mintemperature := 15
		maxtemperature := 30

		for emp := 0; emp < K; emp++ {
			var operator string

			_, err := fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Invalid operator")
				return
			}

			var num int
			_, err2 := fmt.Scan(&num)
			if err2 != nil {
				fmt.Println("Invalid number")
				return
			}

			switch operator {
			case ">=":
				if num >= mintemperature {
					mintemperature = num
				}
			case "<=":
				if num >= mintemperature {
					maxtemperature = num
				}
			default:
				fmt.Printf("Invalid operator")
				return
			}

			if mintemperature > maxtemperature {
				fmt.Println(-1)
				continue
			}

			fmt.Println(mintemperature)
		}
	}
}
