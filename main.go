package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	beeep "github.com/gen2brain/beeep"
)

var ticker *time.Ticker
var cfg config

type pomo struct {
	paused	bool
	counter int
	status	string
}

type config struct {
	workTime int
	breakTime int
}

func initializeModel() pomo {
	return pomo {
		paused: false,
		counter: 0,
		status: "pomodore",
	}
}

func loadConfig() config {
	myConfig := config{
		workTime: 25,
		breakTime: 5,
	}
	var data map[string]interface{}

	plan, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error loading confg", err)
		return myConfig 
	}

	err = json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("Error loading confg", err)
		return myConfig 
	}
	fmt.Println(reflect.TypeOf(data["workTime"]))
	myConfig.workTime = int(data["workTime"].(float64))
	myConfig.breakTime = int(data["breakTime "].(float64))

	fmt.Print("hi")
	fmt.Println("Set work time to ", myConfig.workTime)
	fmt.Println("Set break time to ", myConfig.breakTime)
	return myConfig
}

func (p pomo) Init() tea.Cmd {
	cfg = loadConfig()
	ticker = time.NewTicker(time.Second)
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
		if p.status == "pomodore" && p.counter == cfg.workTime {
			p.status = "break"
			p.counter = 0
			beep("Stop work", "It's time to take a break")
		} else if p.status == "break" && p.counter == cfg.breakTime {
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
