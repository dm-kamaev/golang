//usr/local/bin/go run $0 $@ ; exit

package main

import (
	"fmt"
	// "github.com/pkg/errors"
	"os"
	// "io/ioutil"ig
	// "reflect"
)

func main() {
	// fmt.Println(errors)
	os.Exit(0)
	// ===========================
	// var empty_person Person
	// empty_person2 := new(Person)
	// fmt.Println(empty_person, *empty_person2)

	// ===========================
	person := Person{first_name: `Petya`, last_name: "Ivanov", year: 26}
	// fmt.Println(person)

	// fmt.Println(greet(&person))
	// str := "Hello world"
	// person.say(&str)

	// teacher := Teacher{Person: person, position: "teacher of russin language"}
	// fmt.Println(teacher)

	str2 := "Hello world2"
	teacher := Teacher{Person: person, position: "teacher of russin language"}
	teacher.Person.say(&str2)

	teacher_embeddable := Teacher_embeddable{Person: person, position: "teacher of russin language"}
	teacher_embeddable.say(&str2)

	// fmt.Println(any_argv(5, 2, 3, 4))

	// counter := Counter()
	// fmt.Println(counter(10))
	// fmt.Println(counter(11))

	// fmt.Println(factorial(4))
	// test_read_file()
	// test_defer()
	// test_recover()
	// fmt.Println(task1([]uint{1, 2, 3, 4}))

	// // ===========================
	// fmt.Println(task2(1))
	// fmt.Println(task2(2))
	// fmt.Println(task2(4))
	// fmt.Println(task2(10))
	// fmt.Println(task2(13))
	// // ===========================

	// fmt.Println(search_min(100, 2, 3, 4, 5, -1))

	// // ===========================
	// oddGenerator := makeOddGenerator()
	// fmt.Println(oddGenerator())
	// fmt.Println(oddGenerator())
	// fmt.Println(oddGenerator())
	// fmt.Println(oddGenerator())
	// ===========================
}

type Person struct {
	first_name string
	last_name  string
	year       int
}

func greet(person *Person) string {
	return "Hello, " + person.first_name + " " + person.last_name + "!"
}

func (person *Person) get_full_name() string {
	return person.first_name + " " + person.last_name
}

func (person *Person) say(str *string) {
	fmt.Println(person.get_full_name() + " say: '" + *str + "'")
}

type Teacher struct {
	Person   Person // use teacher.Person.say()
	position string
}

type Teacher_embeddable struct {
	Person   // use teacher.say()
	position string
}

// // =================
// type Error_request struct {
// 	msg  string
// 	Name string
// }

// func new_error_request(msg string) error {
// 	return &Error_request{msg, "Error_request"}
// }

// func (me *Error_request) Error() string {
// 	return me.msg
// }
