package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt1 := intSeq()

	fmt.Println(nextInt1())
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())

	// This initializes a new copy of the closure
	// so we get a new starting value for i
	nextInt2 := intSeq()
	fmt.Println(nextInt2())
}
