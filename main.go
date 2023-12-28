package main

import (
	"fmt"
	"time"
)

func main() {
	onBreak := false

	// TODO: Change this to minutes
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	counter := 0

	go func() {
		for {
			<-ticker.C
			counter++
			fmt.Println("counter at ", counter)
			if !onBreak && counter == 25 {
				onBreak = true
				counter = 0
				fmt.Println("Done work")
				fmt.Print("\x07")
			} else if onBreak && counter == 5 {
				onBreak = false
				counter = 0
				fmt.Println("Done break")
				fmt.Print("\x07")
			}
		}
	} ()

	time.Sleep(time.Minute)
}