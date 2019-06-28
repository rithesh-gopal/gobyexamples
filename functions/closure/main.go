package closure

import "fmt"
import "./lib"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
//Closure is the main class in closure package
func Closure() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	pos, neg:=lib.Adder(),lib. Adder()
	
	for i:=0;i<10;i++{
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}
