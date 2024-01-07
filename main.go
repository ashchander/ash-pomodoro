package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type pomo struct {
	paused	bool
	counter uint8
	status	string
	ticker	*time.Ticker
}

func initializeModel() pomo {
	return pomo {
		paused: false,
		counter: 0,
		status: "pomodore",
		ticker: time.NewTicker(time.Second),
	}
}

func (p pomo) Init() tea.Cmd {
	// No Initial I/O
	return nil
}

func (p pomo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Keypress?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return p, tea.Quit
		case "space", "p":
			p.paused = true
		case "r":
			p.counter = 0
			p.status = "pomodore"
			p.paused = false
		}
	}
	return p, nil
}

func (p pomo) View() string {
	// TODO Implement view code
	return ""
}

func beep(text string) {
	fmt.Println(text)
	fmt.Print("\x07")
}

func handleTicker(ticker *time.Ticker) {
	onBreak := false
	counter = 0
	for {
		<-ticker.C
		if paused {
			continue
		}
		counter++
		// fmt.Println("counter at ", counter)
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
	pomodoro := initializeModel()
	defer pomodoro.ticker.Stop()

	go handleTicker(ticker)

	for text != "quit" {
		if text == "pause" {
			paused = true
		} else if text == "unpause" {
			paused = false
		} else if text == "restart" {
			paused = false
			counter = 0
		}

	}
}
