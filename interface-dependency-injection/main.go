package main

import "log"

// database
type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}

type FamouseDatabase struct{}

func (db FamouseDatabase) GetUser() string {
	return "Will Smith"
}

func (db FamouseDatabase) GetAllUsers() []string {
	return []string{}
}

// greeter
type Greeter interface {
	Greet(userName string)
}
type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s!! Nice to meet you", userName)
}

type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	// db := DefaultDatabase{}
	db := FamouseDatabase{}
	greeter := NiceGreeter{}
	p := Program{
		Db:      db,
		Greeter: greeter,
	}

	p.Execute()

}
