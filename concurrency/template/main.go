package main

import "sync"

func main() {
	// business logics
	rawInputs, _, err := //some api data (as a List)

	var wg sync.WaitGroup
	itemList := []*model.Item{}

	wg.Add(N) // initiate N number of goroutines
	chanOut := make(chan *Item, len(rawInputs))

	for index, item := range rawInputs {
		go func(index int, item any) {
			newItem := // ...
			// business logics
			chanOut <- newItem // assign value to channel
			wg.Done() // remove the goroutine(thead) from the waitGroup
		}(index, item)
	}
	wg.Wait() // wait for all thead to finish
	close(chanOut) // close channel


	for {
		item, open := <- chanOut // return channel output and assign it to a variable
		if !open { break }
		itemsList = append(tagsList, item)
	}

	fmt.Println(itemsList)
}
