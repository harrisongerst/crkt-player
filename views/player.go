package views

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	queue "hgerst/crkt/queuestreamer"
)

var Queue queue.QueueStreamer

type playerModel struct {
	timeElapsed int
	songTitle   string
	queue       *queue.QueueStreamer
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m playerModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Music Player")
}

func PlayerInitialModel() playerModel {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	
	Queue.Current = 0
	speaker.Play(&Queue)

	return playerModel{
		queue:       &Queue,
	}
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated playerModel accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m playerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
		}

	case timeMsg:
		m.timeElapsed++
		if m.timeElapsed > 100 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

// View returns a string based on data in the playerModel. That string which will be
// rendered to the terminal.
func (m playerModel) View() string {
	return fmt.Sprintf("you are listening to %v, it has been %v", m.songTitle, m.timeElapsed)
}

type timeMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return timeMsg{}
}
