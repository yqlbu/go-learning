package main

import (
	"fmt"
	"time"
)

func main() {
	evils := []string{"Tommy", "Johnny", "Bobby", "Andy"}
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	for _, evil := range evils {
		go attack(evil)
	}

	time.Sleep(time.Second * 2)

}

func attack(target string) {
	fmt.Println("Throwing bad guy stars at", target)
}
