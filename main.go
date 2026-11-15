package main

import "fmt"

func main() {
	fmt.Println("=== GO CALCULATOR ===")
	fmt.Println("Welcome!")

	num1 := 20
	num2 := 4

	fmt.Println("First number : ", num1)
	fmt.Println("Second number : ", num2)

	fmt.Println("Add:", num1+num2)
	fmt.Println("Minus:", num1-num2)
	fmt.Println("Divide:", num1/num2)
	fmt.Println("Times:", num1*num2)
}
