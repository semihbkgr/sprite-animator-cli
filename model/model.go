package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintln("press 'q' to exit")
}

func Start() error {
	m := model{}
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}
