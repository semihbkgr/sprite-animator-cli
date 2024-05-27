package model

import (
	"fmt"

	"github.com/semihbkgr/sprite-animator-cli/sprite"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	s sprite.Sprite
}

func new(s sprite.Sprite) model {
	return model{s: s}
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
	s := fmt.Sprintln("press 'q' to exit")
	for _, r := range m.s[0] {
		for _, p := range r {
			s += lipgloss.NewStyle().Foreground(lipgloss.Color(p.ToHexString())).Render("\u2588\u2588")
		}
		s += "\n"
	}
	return s
}

func Start(s sprite.Sprite) error {
	m := new(s)
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}
