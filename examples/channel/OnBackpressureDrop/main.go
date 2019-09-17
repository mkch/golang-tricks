package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	var ch = make(chan int)

	// The sender
	go func() {
		for value := 0; value < 10; value++ {
			select {
			case ch <- value:
				// Sent successfully.
			default:
				// Drop.
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2)))
		}
		close(ch)
	}()

	// The receiver.
	for v := range ch {
		log.Printf("Received: %v\n", v)
	}
}
