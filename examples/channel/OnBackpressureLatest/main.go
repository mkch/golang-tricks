package main

import (
	"log"
	"time"
)

func main() {
	var ch = make(chan int, 1) // Buffer 1

	// The sender
	go func() {
		for value := 0; value < 10; value++ {
			select {
			case <-ch:
				// Discard old value if any.
			default:
				// Nop if no old value.
			}
			// Send the latest.
			ch <- value
		}
		close(ch)
	}()

	time.Sleep(time.Microsecond * 5)
	// The receiver
	for v := range ch {
		// 9 is guaranteed to be received.
		log.Printf("Received: %v", v)
	}
}
