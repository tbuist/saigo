package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	tmp := *y
	*y = *x
	*x = tmp
}

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
