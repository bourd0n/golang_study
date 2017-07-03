package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var result float64 = float64(1)
	var resultPrevious float64 = float64(0)
	for math.Abs(result-resultPrevious) >= 0.00001 {
		fmt.Printf("result %v resultPrev %v\n", result, resultPrevious)
		resultPrevious = result
		result = resultPrevious - (float64(resultPrevious*resultPrevious)-x)/(2.0*resultPrevious)
	}
	return result
}

func main() {
	x := 55.0

	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
