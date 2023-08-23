package main

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	log.Printf("run with %d\n", n)
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup

	// Use the common pool.
	syncCalculateSum := func() {
		time.Sleep(10 * time.Millisecond)
		log.Info().Msg("Hello World!")
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	log.Printf("running goroutines: %d\n", ants.Running())
	log.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		// , ants.WithLogger(&log.Logger)
		wg.Done()
	},
		ants.WithLogger(&log.Logger),
	)
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	log.Printf("running goroutines: %d\n", p.Running())
	log.Printf("finish all tasks, result is %d\n", sum)
}
