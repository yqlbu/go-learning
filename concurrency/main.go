package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	startTime := start()

	timestamps := make([]string, 10)

	// non-parallel execution
	// for i := range timestamps {
	// 	timestamp(i, startTime)
	// }

	// parallel execution
	var wg sync.WaitGroup

	// alternative way to register worker
	wg.Add(len(timestamps))

	for i := range timestamps {
		// copying the value instead of manipulating the pointer
		idx := i

		// parallel execution
		go func() {
			// telling the WG that the exection is done
			defer wg.Done()
			timestamp(idx, startTime)
		}()
	}

	// telling the WG to wait for all the executions to finish
	wg.Wait()

	end(startTime)
}

func timestamp(index int, startTime time.Time) {
	time.Sleep(time.Millisecond * 200)
	t := time.Now().Sub(startTime).String()
	log.Printf("TIMESTAMP %v: %s", index, t)
}

func start() time.Time {
	log.Print("START")
	return time.Now()
}

func end(startTime time.Time) {
	log.Print("END")
	duration := time.Now().Sub(startTime).String()
	log.Printf("FULL DURATION: %s", duration)
}
