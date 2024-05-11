package list

import (
	"github.com/JulianH99/ssshhh/internal/ui"
	l "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	items        []listItem
	view         l.Model
	selectedItem string
}

func New(items []l.Item) List {
	delegate := GetListDelegate()
	list := l.New(items, delegate, 0, 0)

	return List{
		view:         list,
		items:        fromBubbleArray(items),
		selectedItem: "",
	}
}

func (m List) Init() tea.Cmd {
	return nil
}

func (m List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		}

	case tea.WindowSizeMsg:
		m.view.SetSize(msg.Width, msg.Height)
	case ui.CreateKeyMsg:
		return m, m.view.NewStatusMessage("Create new key...")
	}

	m.view, cmd = m.view.Update(msg)
	return m, cmd
}

func (m List) View() string {
	return m.view.View()
}
