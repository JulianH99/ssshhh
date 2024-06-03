package main

import (
	"log"

	"github.com/JulianH99/ssshhh/internal/ui"
	"github.com/JulianH99/ssshhh/internal/ui/create"
	uiList "github.com/JulianH99/ssshhh/internal/ui/list"
	"github.com/JulianH99/ssshhh/internal/ui/modify"
	tea "github.com/charmbracelet/bubbletea"
)

// 1. generate ssh keys with the usual ssh-keygen command
// 2. add it to config file with possibly custom ssh user, domain, etc
// 3. list available keys in use from current config file

type viewState int

const (
	listView viewState = iota
	createKeyView
	modifyKeyView
)

type model struct {
	state  viewState
	list   uiList.List
	create create.Create
	modify modify.Modify
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
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case ui.CreateKeyMsg:
		m.state = createKeyView
	case ui.KeyCreatedMsg:
		m.modify.SetKeyPath(msg.Key)
		m.state = modifyKeyView
		cmds = append(cmds, tea.Cmd(tea.ClearScreen))
	case ui.SshFileEditedMsg:
		newModel, newCmd := m.list.Update(msg)
		m.list = newModel.(uiList.List)
		m.state = listView
		cmds = append(cmds, tea.Cmd(tea.ClearScreen), newCmd)
	}

	switch m.state {
	case listView:
		newModel, newCmd := m.list.Update(msg)
		cmd = newCmd
		m.list = newModel.(uiList.List)
	case createKeyView:
		newModel, newCmd := m.create.Update(msg)
		cmd = newCmd
		cmds = append(cmds, m.create.Init())
		m.create = newModel.(create.Create)
	case modifyKeyView:
		newModel, newCmd := m.modify.Update(msg)
		cmd = newCmd
		cmds = append(cmds, m.modify.Init())
		m.modify = newModel.(modify.Modify)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var view string
	switch m.state {
	case listView:
		view = m.list.View()
	case createKeyView:
		view = m.create.View()
	case modifyKeyView:
		view = m.modify.View()
	default:
		view = m.list.View()
	}

	return ui.AppStyle.Render(view)
	// return view
}

func main() {

	listModel := uiList.New()
	createModel := create.New()
	modifyModel := modify.New()

	initialModel := model{
		state:  listView,
		list:   listModel,
		create: createModel,
		modify: modifyModel,
	}

	p := tea.NewProgram(initialModel)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
