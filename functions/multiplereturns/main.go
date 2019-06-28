package multiplereturns

import "fmt"

func vals() (int, int) {
	return 2, 3
}
//MultipleReturn is the main method in multiplereturns package
func MultipleReturn() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
