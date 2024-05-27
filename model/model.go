package model

import (
	"fmt"
	"time"

	"github.com/semihbkgr/sprite-animator-cli/sprite"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	s sprite.Sprite
	i int
}

func new(s sprite.Sprite) model {
	return model{s: s, i: 0}
}

type AnimationTimer struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(_ time.Time) tea.Msg {
		return AnimationTimer{}
	})
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case AnimationTimer:
		m.i++
		if m.i == len(m.s) {
			m.i = 0
		}
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintln("press 'q' to exit")
	s += fmt.Sprintf("%d/%d\n", m.i, len(m.s))
	for _, r := range m.s[m.i] {
		for _, p := range r {
			if p.IsTransparent() {
				s += "  "
			} else {
				s += lipgloss.NewStyle().Foreground(lipgloss.Color(p.ToRGBHexString())).Render("\u2588\u2588")
			}
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
