package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func beep(text string) {
	fmt.Println(text)
	fmt.Print("\x07")
}

func readInput(reader *bufio.Reader) (string, error)  {
	fmt.Print("Enter a command: ")
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text, err
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

	reader := bufio.NewReader(os.Stdin)
	var text string
	var err error
	for text != "quit" {
		text, err = readInput(reader)
		if err != nil {
			log.Fatal("Could not read input")
		}
	}
}
