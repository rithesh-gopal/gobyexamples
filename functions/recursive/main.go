package recursive

import "fmt"

func fact(n int)int{
	if n==0{
		return 1
	}
	return n* fact(n-1)
}
//Recursion main
func Recursion(){
	fmt.Println(fact(7))
}