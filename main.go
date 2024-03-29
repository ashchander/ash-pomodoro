package main

import (
	"fmt"
	"strconv"
	"time"
	tea "github.com/charmbracelet/bubbletea"
	beeep "github.com/gen2brain/beeep"
)

var ticker *time.Ticker

type pomo struct {
	paused	bool
	counter int
	status	string
}

func initializeModel() pomo {
	return pomo {
		paused: false,
		counter: 0,
		status: "pomodore",
	}
}

func (p pomo) Init() tea.Cmd {
	ticker = time.NewTicker(time.Minute)
	return tick
}

func (p pomo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Keypress?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return p, tea.Quit
		case "space", "p":
			p.paused = true
		case "r":
			p.counter = 0
			p.status = "pomodore"
			p.paused = false
		}
	// Ticker?
	case tickMsg:
		if p.paused {
			return p, tick	
		}
		p.counter++
		if p.status == "pomodore" && p.counter == 25 {
			p.status = "break"
			p.counter = 0
			beep("Stop work", "It's time to take a break")
		} else if p.status == "break" && p.counter == 5 {
			p.status = "pomodore"
			p.counter = 0
			beep("Break over", "It's time to get back to work")
		}
	}

	return p, tick
}

func (p pomo) View() string {
	// TODO make UI nicer
	var s string
	if p.status == "pomodore" {
		s = "Focus time!!!!\n"
	} else {
		s = "Break time!!!!\n"
	}
	s += "Currently on minute: " + strconv.Itoa(p.counter + 1)
	return s
}

func beep(title, text string) {
	fmt.Print("\x07")
	err := beeep.Notify(title, text, "assets/information.png")
	if err != nil {
	    panic(err)
	}
}

type tickMsg time.Time

func tick() tea.Msg {
	// Send Tick Message for each ticker event
	<- ticker.C
	return tickMsg{}
}

func main() {
	pomodoro := initializeModel()
	p := tea.NewProgram(pomodoro)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
