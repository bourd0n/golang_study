package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var result float64 = float64(1)
	var resultPrevious float64 = float64(0)
	for math.Abs(result-resultPrevious) >= 0.01 {
		fmt.Printf("result %v resultPrev %v\n", result, resultPrevious)
		resultPrevious = result
		result = resultPrevious - (float64(resultPrevious*resultPrevious)-x)/(2.0*resultPrevious)
	}
	return result, nil
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
