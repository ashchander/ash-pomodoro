package main

import (
	"fmt"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type pomo struct {
	paused	bool
	counter int
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
	// Send message on ticks
	return tick
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
	// Ticker?
	case tickMsg:
		if p.paused {
			return p, nil	
		}
		p.counter++
		if p.status == "pomodore" && p.counter == 25 {
			p.status = "break"
			p.counter = 0
			beep("Done work")
		} else if p.status == "break" && p.counter == 5 {
			p.status = "pomodore"
			p.counter = 0
			beep("Done with break")
		}
	}

	return p, nil
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

func beep(text string) {
	fmt.Println(text)
	fmt.Print("\x07")
}

type tickMsg time.Time

func tick() tea.Msg {
	// TODO change to Minutes
	time.Sleep(time.Second)
	return tickMsg{}
}

func main() {
	pomodoro := initializeModel()
	defer pomodoro.ticker.Stop()
	p := tea.NewProgram(pomodoro)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
