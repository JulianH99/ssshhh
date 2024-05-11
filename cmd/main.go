package main

import (
	"log"

	"github.com/JulianH99/ssshhh/internal/config"
	uiList "github.com/JulianH99/ssshhh/internal/ui/list"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// 1. generate ssh keys with the usual ssh-keygen command
// 2. add it to config file with possibly custom ssh user, domain, etc
// 3. list available keys in use from current config file

type viewState int

const (
	listView viewState = iota
	createKeyView
)

type model struct {
	state viewState
	list  uiList.List
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.state == listView {
		newModel, newCmd := m.list.Update(msg)
		cmd = newCmd
		m.list = newModel.(uiList.List)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.state {
	case listView:
		return m.list.View()
	default:
		return m.list.View()
	}
}

func createKeyForm() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("path").
				Title("Path").
				Placeholder("~/.ssh/id_ed25519"),
			huh.NewSelect[string]().
				Key("type").
				Title("Encryption key").
				Options(
					huh.NewOption("ed25519", "ed25519"),
					huh.NewOption("rsa", "rsa"),
				),
			huh.NewInput().
				Key("comment").
				Title("Comment").
				Placeholder("user@hostname.com").
				Description("Enter a comment for your key"),
		),
	)

	return form

}

func configToListItems(sshConfigs []config.SshConfig) []list.Item {

	listItems := make([]list.Item, len(sshConfigs))

	for i, sshConfig := range sshConfigs {
		listItems[i] = uiList.NewItem(sshConfig.Host, sshConfig.User)
	}

	return listItems

}

func main() {
	sshConfigs, err := config.SshConfigs()

	if err != nil {
		log.Fatal(err)
	}

	items := configToListItems(sshConfigs)

	listModel := uiList.New(items)

	initialModel := model{
		state: listView,
		list:  listModel,
	}

	p := tea.NewProgram(initialModel)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
