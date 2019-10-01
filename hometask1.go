package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Print("Enter x: ")
	var x float64
	fmt.Scanf("%f", &x)

	fmt.Print("Enter y: ")

	var y float64
	fmt.Scanf("%f", &y)

	fmt.Print("Enter z: ")
	var Z string
	fmt.Scanf("%s", &z)
	fmt.Println(Z)

}

func calc() {

	switch z {
	case "*":
		fmt.Println(x * y)
	case "-":
		fmt.Println(x - y)
	case "+":
		fmt.Println(x + y)
	case "/":
		fmt.Println(x / y)
	default:
		fmt.Println("err")
	}

}
