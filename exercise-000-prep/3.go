package main

import (
	"fmt"
)

func half(x int) (float64, bool) {
	return float64(x) / float64(2), x%2 == 0
}

func half_test(in int, exp1 float64, exp2 bool) {
	res1, res2 := half(in)
	fmt.Println(res1, res2, "should equal", exp1, exp2)
}

func greatest(xs ...int) (greatest int) {
	greatest := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > greatest {
			greatest = xs[i]
		}
	}
	return
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	half_test(7, 3.5, false)
	half_test(8, 4, true)
	half_test(1, 0.5, false)
	half_test(2, 1, true)
}
