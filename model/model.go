package model

import (
	"time"

	"github.com/semihbkgr/sprite-animator-cli/sprite"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	s     sprite.Sprite
	i     int
	start int
	end   int
	fps   int
}

func new(s sprite.Sprite, start int, end int, fps int) model {
	return model{s: s, i: start, start: start, end: end, fps: fps}
}

type AnimationTimer struct{}

func tick(fps int) tea.Cmd {
	return tea.Tick(time.Duration(time.Second.Nanoseconds()/int64(fps)), func(_ time.Time) tea.Msg {
		return AnimationTimer{}
	})
}

func (m model) Init() tea.Cmd {
	return tick(m.fps)
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case AnimationTimer:
		m.i++
		if m.i > m.end {
			m.i = m.start
		}
		return m, tick(m.fps)
	}

	return m, nil
}

func (m model) View() string {
	s := ""
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

func Start(s sprite.Sprite, start, end, fps int) error {
	m := new(s, start, end, fps)
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}
