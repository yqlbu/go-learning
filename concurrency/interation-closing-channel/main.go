package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	go throwingNinjaStar(channel)
	for {
		message, open := <-channel
		// check if the channel is still open
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3

	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("Your scored: ", score)
	}

	// close channel, when the above is done
	close(channel)
}
