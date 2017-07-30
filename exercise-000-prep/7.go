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

func greatest(xs ...float64) (great float64) {
	great = xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > great {
			great = xs[i]
		}
	}
	return
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

func fib_rec(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib_rec(n-1) + fib_rec(n-2)
	}
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	half_test(7, 3.5, false)
	half_test(8, 4, true)
	half_test(1, 0.5, false)
	half_test(2, 1, true)
	fmt.Println(greatest(xs...))
	gen := makeOddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
	for i := 0; i < 10; i++ {
		fmt.Println(fib_rec(i))
	}
}
