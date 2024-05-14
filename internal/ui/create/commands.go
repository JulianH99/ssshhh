package create

import (
	tea "github.com/charmbracelet/bubbletea"
)

type execSshKeygen struct{}

func execSshKeygenCmd() tea.Msg {
	return execSshKeygen{}
}
