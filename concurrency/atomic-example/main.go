package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	var m sync.Mutex
	m.Lock()
	sum++
	m.Unlock()
	fmt.Println(sum)

	var diffSum int64
	fmt.Println(atomic.LoadInt64(&diffSum))
	atomic.StoreInt64(&diffSum, 1)
	fmt.Println(diffSum)

	var av atomic.Value
	wallace := ninja{name: "wallace"}
	av.Store(wallace)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		w := av.Load().(ninja)
		w.name = "Not Wallace"
		// pass in a copy of w object
		av.Store(w)
		w.name = "Wallace Again"
		wg.Done()

	}()
	wg.Wait()
	fmt.Println(av.Load().(ninja).name)
}

type ninja struct {
	name string
}
