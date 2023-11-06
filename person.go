package testlib

import "fmt"

var a = person{name: "ake", age: 12}
var b = person{name: "bath", age: 15}

type person struct {
	name string
	age int
}

func init() {
	fmt.Println(a.age)
	fmt.Println(b.name)
}

func PersonAge(id string) int{
	if id =="a" {
		return a.age
	}
	if id == "b" {
		return b.age
	}
	return 0
}

