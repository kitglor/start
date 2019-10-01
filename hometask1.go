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
	var z string
	fmt.Scanf("%s", &z)

	fmt.Println(kalk(z))
}
func kalk(x float64, y float64, z string) {
	{

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
			return z
		}

	}
}
