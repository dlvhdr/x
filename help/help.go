package help

import (
	helpBubble "charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	bHelp       helpBubble.Model
	width       int
	keys        [][]key.Binding
	Styles      Styles
	searchInput textinput.Model
}

type Styles struct {
	helpBubble.Styles
	RootStyle lipgloss.Style
}

func New() Model {
	m := Model{}
	m.bHelp = helpBubble.New()
	helpSt := lipgloss.NewStyle()
	m.bHelp.ShortSeparator = " · "
	m.bHelp.Styles.FullKey = helpSt
	m.bHelp.Styles.FullDesc = helpSt
	m.bHelp.Styles.FullSeparator = helpSt
	m.bHelp.Styles.FullKey = helpSt.Foreground(lipgloss.Blue)
	m.bHelp.Styles.FullDesc = helpSt
	m.bHelp.Styles.FullSeparator = helpSt
	m.bHelp.Styles.Ellipsis = helpSt
	m.Styles.RootStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Padding(1, 3).BorderForeground(lipgloss.Blue)

	m.searchInput = textinput.New()

	return m
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetWidth(msg.Width / 2)
	}

	return m, nil
}

func (m *Model) Width() int {
	return m.width
}

func (m *Model) SetWidth(width int) {
	m.width = width
}

func (m *Model) SetKeys(groups [][]key.Binding) {
	m.keys = groups
}

func (m *Model) View() string {
	m.updateStyles()

	helpContent := m.bHelp.FullHelpView(m.keys)

	return m.Styles.RootStyle.Render(helpContent)
}

func (m *Model) updateStyles() {
	m.bHelp.Styles.FullKey = m.Styles.FullKey
	m.bHelp.Styles.FullDesc = m.Styles.FullDesc
	m.bHelp.Styles.FullSeparator = m.Styles.FullSeparator
	m.bHelp.Styles.FullKey = m.Styles.FullKey
	m.bHelp.Styles.FullDesc = m.Styles.FullDesc
	m.bHelp.Styles.FullSeparator = m.Styles.FullSeparator
	m.bHelp.Styles.Ellipsis = m.Styles.Ellipsis
}
