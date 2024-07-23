package views

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

func SelectInitialModel() model {
	return model{
		choices: listFilesInDirectory(),
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("file list")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}	
			Queue.Add(createStreamer(m.choices[m.cursor]))
		case "p", "g":
				p := tea.NewProgram(PlayerInitialModel())
				if _, err := p.Run(); err != nil {
				fmt.Printf("error running player: %v", err)
				os.Exit(1)
			}
		}
		
	}

	return m, nil
}

func (m model) View() string {
	s := "what song do you want to play \n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func listFilesInDirectory() []string {
	files, err := os.ReadDir("sounds/")
	if err != nil {
		fmt.Printf("error reading directory: %v", err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

//implement
func createStreamer(filePath string) beep.StreamSeekCloser {
	f, err := os.Open("sounds/" + filePath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return streamer
}