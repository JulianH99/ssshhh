package modify

import (
	"github.com/JulianH99/ssshhh/internal/config"
	"github.com/JulianH99/ssshhh/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func editSshConfig(sshConfig config.SshConfig) tea.Cmd {
	return func() tea.Msg {
		err := config.AddEntryToSshConfig(sshConfig)

		if err != nil {
			return ui.ErrorMsg{Err: err}
		}

		return ui.SshFileEditedMsg{}
	}
}
