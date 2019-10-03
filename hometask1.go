package main

import (
	"fmt"
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

	fmt.Println(kalk(x, y, z))
}
func kalk(x float64, y float64, z string) float64 {
	{

		switch z {
		case "*":
			return (x * y)
		case "-":
			return (x - y)
		case "+":
			return (x + y)
		case "/":
			return (x / y)
		default:
			return (0)

		}
	}
}
