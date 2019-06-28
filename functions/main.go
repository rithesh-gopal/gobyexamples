package main

import "fmt"
import C "./closure"
import M "./multiplereturns"
import V "./variadic"
import R "./recursive"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	fmt.Println("Closures......")
	C.Closure()

	fmt.Println("MultipleReturn......")
	M.MultipleReturn()

	fmt.Println("Variadic......")
	V.Variadic()
	
	fmt.Println("Recursive...")
	R.Recursion()
}
