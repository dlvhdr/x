package main

import (
	"fmt"
	"log"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/dlvhdr/x/help"
)

var exampleKeys = []key.Binding{
	key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "expand"),
	),
	key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "collapse"),
	),
	key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "toggle"),
	),
	key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "prev file"),
	),
	key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "next file"),
	),
	key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "diff down"),
	),
	key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("ctrl+u", "diff up"),
	),
	key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "toggle file tree"),
	),
	key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "search files"),
	),
	key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "copy file path"),
	),
	key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch panel"),
	),
	key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open"),
	),
	key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "toggle side-by-side"),
	),
	key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "toggle icon style"),
	),
	key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}

type programModel struct {
	h *help.Model
}

func (m programModel) Init() tea.Cmd {
	return nil
}

func (m programModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		}
	}

	_, cmd := m.h.Update(msg)
	return m, cmd
}

func (m programModel) View() tea.View {
	var v tea.View
	v.SetContent(m.h.View())
	return v
}

func main() {
	h := help.New()
	h.Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Red)
	h.Styles.RootStyle = h.Styles.RootStyle.BorderForeground(lipgloss.Red)
	h.SetKeys([][]key.Binding{exampleKeys})

	p := tea.NewProgram(programModel{h: &h})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(h.View())
}
