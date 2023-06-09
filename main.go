package main

import (
	"fmt"
	"module_name/calc"
)

func main() {
	fmt.Println(calc.Add(4, 5))
	fmt.Println(calc.Div(4, 5))
	fmt.Println(calc.Sub(4, 5))
	fmt.Println(calc.Mult(4, 5))
}
