package main

import "fmt"

func main() {

	fmt.Print("Enter x: ")
	var x float64
	fmt.Scanf("%f", &x)

	fmt.Print("Enter y: ")

	var y float64
	fmt.Scanf("%f", &y)

	fmt.Print("Enter z: ")
	var z string
	fmt.Scanf("%s", &z)
	fmt.Println(z)

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
