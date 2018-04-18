//usr/local/bin/go run $0 $@ ; exit

package main

import (
	"fmt"
	// "io/ioutil"
	// "reflect"
)

func main() {
	// ===========================
	number := 23
	res := test_pointer(&number)
	fmt.Println(*res)

	// ===========================
	number_1 := 12
	res2 := create_new_pointer(&number_1)
	fmt.Println(number_1, res2)

	// ===========================
	num := 1.5
	square(&num)
	fmt.Println(num)

	// ===========================
	x := 1
	y := 2
	swap_1(&x, &y)
	fmt.Println(x, y)

	// ===========================
	z := 1
	f := 2
	swap_2(&z, &f)
	fmt.Println(z, f)

}

func test_pointer(number *int) *int {
	*number = *number * 2
	return number
}

func create_new_pointer(number *int) int {
	ref := new(int)
	*ref = *number * 2
	return *ref
}

func square(x *float64) {
	*x = *x * *x
}

func swap_1(x *int, y *int) {
	temp := new(int)
	*temp = *y
	*y = *x
	*x = *temp
}

func swap_2(x *int, y *int) {
	*x, *y = *y, *x
}
