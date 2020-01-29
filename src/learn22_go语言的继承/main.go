package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) say(content string) {
	fmt.Println(a.name, " say ", content)
}

type Dog struct {
	*Animal
}

func newDog(name string, age int) *Dog {
	return &Dog{&Animal{name: name, age: age}}
}
func main() {
	d := newDog("王兰花", 18)
	d.say("hello go lang")
}
