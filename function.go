//usr/local/bin/go run $0 $@ ; exit

package main

import (
	"fmt"
	"io/ioutil"
	// "reflect"
)

func main() {
	res := average([5]float64{1, 2, 3, 4, 5})
	fmt.Println(res)

	fmt.Println(retun_many())

	fmt.Println(many_ints(1, 2, 3, 4))

	fmt.Println(any_argv(5, 2, 3, 4))

	counter := Counter()
	fmt.Println(counter(10))
	fmt.Println(counter(11))

	fmt.Println(factorial(4))
	test_read_file()
	test_defer()
	test_recover()
	fmt.Println(task1([]uint{1, 2, 3, 4}))

	// ===========================
	fmt.Println(task2(1))
	fmt.Println(task2(2))
	fmt.Println(task2(4))
	fmt.Println(task2(10))
	fmt.Println(task2(13))
	// ===========================

	fmt.Println(search_min(100, 2, 3, 4, 5, -1))

	// ===========================
	oddGenerator := makeOddGenerator()
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	// ===========================
}

func average(ar [5]float64) float64 {
	total := 0.0
	for _, val := range ar {
		total += val
	}
	return total / float64(len(ar))
}

func retun_many() ([]int, []int) {
	x := []int{1, 2}
	y := []int{2, 3}
	return x, y
}

func many_ints(arg ...int) int {
	sum := 0
	for _, val := range arg {
		sum += val
	}
	return sum
}

func any_argv(arg ...interface{}) int {
	sum := 0
	for _, val := range arg {
		sum += val.(int) // convert interface to int
	}
	return sum
}

func Counter() func(int) int {
	i := 0
	return func(add int) int {
		i = (i + 1) + add
		return i
	}
}

func factorial(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func test_read_file() {
	data, err := ioutil.ReadFile("./test.txt")
	fmt.Println(string(data), err)
}

func test_defer() {

	first := func() {
		fmt.Println("first")
	}

	second := func() {
		fmt.Println("second")
	}

	// always call before end
	defer second()
	// panic("ERROR")
	first()
}

func test_recover() {
	defer func() {
		str := recover()
		fmt.Println("ERROR= " + str.(string))
	}()

	panic("!!!OOpps ERROR!!!")
}

//Функция sum принимает срез чисел и складывает их вместе. Как бы выглядела сигнатура этой функции?
func task1(list []uint) uint {
	var total uint = 0
	for _, val := range list {
		total += val
	}
	return total
}

// Напишите функцию, которая принимает число,
// делит его пополам и возвращает true в случае, если получившееся число чётное, и false в случае нечетного результата.
// Например, half(1) должна вернуть (0, false), в то время как half(2) вернет (1, true).
func task2(num float64) (int, bool) {
	half := int(num / 2)

	if half == 0 {
		return 0, false
	}

	is_even := (half % 2) == 0
	if is_even {
		return 1, true
	} else {
		return 0, false
	}
}

func search_min(argv ...int) int {
	var min int
	for _, val := range argv {
		if val < min {
			min = val
		}
	}
	return min
}

func makeOddGenerator() func() int {
	var i int
	return func() int {
		if i == 0 {
			i++
			return 0
		} else {
			if i%2 != 0 {
				old := i
				i++
				return old
			} else {
				i++
				old := i
				i++
				return old
			}
		}
	}
}
