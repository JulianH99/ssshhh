package list

import (
	"log"

	"github.com/JulianH99/ssshhh/internal/config"
	"github.com/JulianH99/ssshhh/internal/ui"
	l "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	items        []listItem
	view         l.Model
	selectedItem string
}

func New() List {
	delegate := GetListDelegate()
	items := readSshConfigItems()
	list := l.New(items, delegate, 0, 0)
	list.Styles.Title = ui.ListTitleStyle
	list.Styles.StatusBar = ui.ListStatusBarStyle
	list.Title = "Available ssh configurations"

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
		m.view.SetSize(msg.Width, msg.Height-2)
	case ui.CreateKeyMsg:
		return m, m.view.NewStatusMessage(ui.StatusMessageStyle("Create new key..."))
	case ui.SshFileEditedMsg:
		items := readSshConfigItems()
		m.items = fromBubbleArray(items)
		m.view.SetItems(items)
		return m, m.view.NewStatusMessage(ui.StatusMessageStyle("Added ssh config entry to file"))
	}

	m.view, cmd = m.view.Update(msg)
	return m, cmd
}

func (m List) View() string {
	return m.view.View()
}

func configToListItems(sshConfigs []config.SshConfig) []l.Item {
	listItems := make([]l.Item, len(sshConfigs))

	for i, sshConfig := range sshConfigs {
		listItems[i] = NewItem(sshConfig.Host, sshConfig.User)
	}

	return listItems
}

func readSshConfigItems() []l.Item {
	sshConfigs, err := config.SshConfigs()

	if err != nil {
		log.Fatal(err)
	}

	return configToListItems(sshConfigs)
}
