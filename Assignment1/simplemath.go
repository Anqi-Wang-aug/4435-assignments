/* input format: [operation][space][number1][space][number2]
 * Examples: add 1 2 = 1+2
 * 			 sub 4 2 = 4-2
 *			 mult 2 3 = 2*3
 * 			 div 4 2 = 4/2
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cmd string
	var num1, num2, ans float64

	//Getting input from users
	cmd = os.Args[1]
	num1, err1 := strconv.ParseFloat(os.Args[2], 64)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil {
		fmt.Println("Invalid input: first number")
		os.Exit(1)
	}

	if err2 != nil {
		fmt.Println("invalid input: second number")
		os.Exit(1)
	}

	switch {
	case (cmd == "add"):
		ans = num1 + num2
		fmt.Println(num1, "+", num2, "=", ans)
	case (cmd == "sub"):
		ans = num1 - num2
		fmt.Println(num1, "-", num2, "=", ans)
	case (cmd == "mult"):
		ans = num1 * num2
		fmt.Println(num1, "*", num2, "=", ans)
	case (cmd == "div"):
		ans = num1 / num2
		fmt.Println(num1, "/", num2, "=", ans)
	default:
		fmt.Println("Oops...an error occurred")
	}
}
