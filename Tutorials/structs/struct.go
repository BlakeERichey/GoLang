package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

//recieves copy of person and simply references the values
func (x person) getAge() int {
	return x.age
}

func (x person) getName() string {
	return x.name
}

//takes pointer so it can modify the original
func (x *person) setAge(newAge int) {
	x.age = newAge
}

func main(){
	fmt.Println(person{name: "Blake"})

	fmt.Println(person{name: "Alice", age: 30})

	s:= person{name: "Alice", age: 30}
	s.name = "James"
	fmt.Println(s)
	fmt.Println(&s)

	fmt.Println(s.getAge())
	s.setAge(15)
	fmt.Println(s)

	fmt.Println(substring(s.getName()))

	arr := createArr(s.getName(), "Adam")
	fmt.Println(arr[])
}

//returns all but the first character of a word
func substring(value string) string {
	return value[1:]
}

//return array of variable size
func createArr(x, y string) []string{
	array := make([]string, 2)
	array[0], array[1] = x, y
	return array
}