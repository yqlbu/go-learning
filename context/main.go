package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	log.Println(ctx.Value("myKey"))
	ctx2 := context.WithValue(ctx, "myKey", 123)

	log.Println(ctx2)
	log.Println(ctx2.Value("myKey"))

	t := time.Now().Add(time.Second * 1)
	ctx3, cancelFn := context.WithDeadline(ctx, t)
	deadline, ok := ctx3.Deadline()
	log.Println(deadline)
	log.Println(ok)
	cancelFn()
	log.Println(ctx3.Err())
	time.Sleep(time.Second * 2)
	log.Println(ctx3.Err() == context.DeadlineExceeded)

	// timeout := time.Second * 1
	// ctx4, _ := context.WithTimeout(ctx, timeout)

	// log.Println(ctx4.Err())
	// time.Sleep(time.Second * 2)
	// log.Println(ctx4.Err() == context.DeadlineExceeded)
	fn(ctx3)
}

func fn(ctx context.Context) {
	select {
	case <-time.After(time.Second * 2):
		log.Print("after 2 sec")
	case <-time.After(time.Second * 1):
		log.Print("after 1 sec")
	case <-ctx.Done():
		log.Println("boom context is done")
	}
}
