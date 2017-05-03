package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second * 1)
	boom := time.After(time.Second * 4)

	for {
		select {
		case <-tick:
			fmt.Println("tick...")
		case <-boom:
			fmt.Println("boom...")
			return
		default:
			fmt.Println("default...")
			time.Sleep(time.Second * 3)
		}
	}
}
