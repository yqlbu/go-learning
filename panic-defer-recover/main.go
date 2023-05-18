package main

import "log"

func main() {
	// -- panic --
	// return non-zero status
	// use panic if we want to exit out the program
	// panic("boom")

	// -- defer --
	// defer will be executed at the end of a function
	// use defer when a goroutine finishes its execution (e.g closing a connection)
	// log.Println("start")

	// defer func() {
	//     log.Println("defer")
	// }

	// log.Println("end")

	// -- recover --
	// output: defer, recovered
	// similar usecase of try catch block in nodejs
	defer func() {
		log.Println("defer")
		if r := recover(); r != nil {
			log.Println("recovered")
		}
	}()

	panic("boom")
	// not executed
	log.Println("carry on")
}
