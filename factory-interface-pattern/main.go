package main

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p *person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// Here, NewPerson returns an interface, and not the person struct itself
func NewPerson(name string, age int) Person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p := NewPerson("Kevin", 16)
	p.Greet()
}
