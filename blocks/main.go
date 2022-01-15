package main

import "fmt"

func main() {
	//
	a, _ := fmt.Scan()
	arithmetics(a, 2)
}

func arithmetics(num1, num2 int) (a, m, d, s int) {

	return num1 + num2, num2 * num1, num1 / num2, num2 - num1
}
