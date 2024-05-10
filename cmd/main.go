package main

import (
	"log"

	"github.com/JulianH99/ssshhh/internal/config"
	"github.com/JulianH99/ssshhh/internal/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// 1. generate ssh keys with the usual ssh-keygen command
// 2. add it to config file with possibly custom ssh user, domain, etc
// 3. list available keys in use from current config file

type item struct {
	title string
	desc  string
}

func (i item) FilterValue() string {
	return i.title
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.desc
}

type viewState int

const (
	listView viewState = iota
	createKeyView
)

type model struct {
	state         viewState
	list          list.Model
	createKeyForm *huh.Form
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)

	case ui.CreateKeyMsg:
		m.state = createKeyView
	}

	if m.state == createKeyView {

		newFormModel, newCmd := m.createKeyForm.Update(msg)

		m.createKeyForm = newFormModel.(*huh.Form)
		cmd = newCmd
		cmds = append(cmds, m.createKeyForm.Init())

		if m.createKeyForm.State == huh.StateCompleted {
			m.state = listView
			cmds = append(cmds, m.list.NewStatusMessage("Key created"))
		}

	} else {
		m.list, cmd = m.list.Update(msg)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.state {
	case listView:
		return m.list.View()
	case createKeyView:

		return m.createKeyForm.View()
	default:
		return m.list.View()
	}
}

func configToListItems(configs []config.SshConfig) []list.Item {

	items := make([]list.Item, len(configs))

	for i, config := range configs {
		items[i] = item{title: config.User, desc: config.Host}
	}

	return items
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

func main() {
	sshConfigs, err := config.SshConfigs()

	if err != nil {
		log.Fatal(err)
	}

	items := configToListItems(sshConfigs)
	delegate := ui.GetListDelegate()

	list := list.New(items, delegate, 0, 0)
	list.Title = "Available SSH configurations"

	initialModel := model{
		state:         listView,
		list:          list,
		createKeyForm: createKeyForm(),
	}
	p := tea.NewProgram(initialModel)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
