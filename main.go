package main

import (
	"fmt"
	"time"
)

func beep(text string) {
	fmt.Println(text)
	fmt.Print("\x07")
}

func handleTicker(ticker *time.Ticker) {
	onBreak := false
	counter := 0
	for {
		<-ticker.C
		counter++
		fmt.Println("counter at ", counter)
		if !onBreak && counter == 25 {
			onBreak = true
			counter = 0
			beep("Done work")
		} else if onBreak && counter == 5 {
			onBreak = false
			counter = 0
			beep("Done with break")
		}
	}
}

func main() {
	// TODO: Change this to minutes
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go handleTicker(ticker)

	// TODO: stay alive until interrupt
	time.Sleep(time.Minute)
}