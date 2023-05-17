// main.go
package main

import "log"

// interface
type Animal interface {
	Sound() string
}

// Cat
type Cat struct {
}

func NewCat() *Cat {
	return &Cat{}
}

func (c Cat) Sound() string {
	return "meow"
}

// Dog
type Dog struct {
	name string
}

func NewDog() *Dog {
	return &Dog{name: "Dough"}
}

func (c Dog) Sound() string {
	return "wooh"
}

func farm(x int) {
	var a Animal
	if x > 42 {
		a = Cat{}
	} else {
		a = Dog{}
	}

	log.Print(a.Sound())
}

func main() {
	farm(52)
}
