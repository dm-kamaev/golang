package main

import "fmt"
import "a"

// import "github.com/go-errors/errors"

func main() {
	// var Crashed = errors.Errorf("oh dear")
	fmt.Println("Hello world")
	greet := a.Greet("Vasya")
	fmt.Println(greet)
	// err := errors.New(Crashed)
	// fmt.Println(err.ErrorStack())
	fmt.Println(2)
}
