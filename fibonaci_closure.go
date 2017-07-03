package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fn1 := 1
	fn2 := 0
	firstCall := true
	return func() int {
		if firstCall {
			firstCall = false
			return 0
		}
		result := fn2 + fn1
		fn1 = fn2
		fn2 = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
