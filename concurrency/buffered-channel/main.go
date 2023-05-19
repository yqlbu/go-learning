package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)
	channel <- "Frist message"
	channel <- "Second message"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

// channel has a FIFO(First In First Out) concept as a queue
