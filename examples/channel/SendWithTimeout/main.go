package main

import (
	"log"
	"time"
)

func main() {
	var ch = make(chan int)
	const timeout = time.Millisecond * 500

	// The sender
	go func() {
		for value := 0; value < 10; value++ {
			timer := time.NewTimer(timeout)
			select {
			case ch <- value:
				// Sent successfully.
			case <-timer.C:
				// Timed out.
				log.Printf("Sending %v failed. Timeout.", value)
			}
			timer.Stop()
		}
		close(ch)
	}()

	// The receiver
	for i := range ch {
		log.Printf("Received %v", i)
		if i == 5 {
			time.Sleep(timeout + time.Millisecond*100)
		}
	}

}
