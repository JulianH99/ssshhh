package list

import (
	"fmt"

	"github.com/JulianH99/ssshhh/internal/ui"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	Create key.Binding
	Choose key.Binding
}

var defaultKeyMap = keymap{
	Create: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Create"),
	),
	Choose: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Choose"),
	),
}

func GetListDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()

	delegate.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i := m.SelectedItem(); i != nil {
			title = i.FilterValue()
		} else {
			title = ""
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, defaultKeyMap.Choose):
				return m.NewStatusMessage(ui.StatusMessageStyle(fmt.Sprintf("Selected %s", title)))
			case key.Matches(msg, defaultKeyMap.Create):
				// need to create modal
				return switchToCreateKeyView
			}
		}

		return nil
	}

	help := []key.Binding{defaultKeyMap.Create, defaultKeyMap.Choose}

	delegate.ShortHelpFunc = func() []key.Binding {
		return help
	}

	delegate.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return delegate

}
