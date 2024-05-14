package list

import (
	"github.com/JulianH99/ssshhh/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func switchToCreateKeyView() tea.Msg {
	return ui.CreateKeyMsg{}
}
