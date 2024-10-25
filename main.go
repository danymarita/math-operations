package main

import (
	"fmt"

	"github.com/danymaritaocean/math-operations/functions"
)

func main() {
	x := 2
	y := 3
	additionResult := functions.Addition(x, y)
	fmt.Println("Addition result: ", additionResult)
	substractionResult := functions.Substraction(x, y)
	fmt.Println("Substraction result: ", substractionResult)
}
