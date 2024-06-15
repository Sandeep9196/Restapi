package main

import (
	"fmt"
	"time"
)

func main() {
	endRoutine := make(chan bool)

	for i := 0; i < 10; i++ {
		routineID := i
		go func() {
			fmt.Printf("Started %d\n", routineID)
			for {
				select {
				case <-endRoutine:
					fmt.Printf("Ended %d\n", routineID)
					return
				default:
					break
				}
				//Do work
				time.Sleep(10)
				fmt.Printf("Still Alive %d\n", routineID)
			}

		}()
	}
	time.Sleep(1)
	endRoutine <- true
	time.Sleep(1)
	endRoutine <- true
	time.Sleep(1)
	endRoutine <- true
	time.Sleep(1)
	endRoutine <- true
}
