package modify

import (
	"github.com/JulianH99/ssshhh/internal/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Modify struct {
	form         *huh.Form
	formFinished bool
	keyPath      string
}

func (m *Modify) SetKeyPath(key string) {
	m.keyPath = key
}

func (m Modify) Init() tea.Cmd {
	return m.form.Init()
}

func (m Modify) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newForm, cmd := m.form.Update(msg)
	m.form = newForm.(*huh.Form)

	if m.form.State == huh.StateCompleted && !m.formFinished {
		m.formFinished = true

		sshConfig := config.CreateNewConfig(
			m.form.GetString("host"),
			m.form.GetString("user"),
			m.form.GetString("hostname"),
			m.keyPath,
			stringToMap(m.form.GetString("extras")),
		)
		return m, editSshConfig(sshConfig)
	}
	return m, cmd
}

func (m Modify) View() string {
	return m.form.View()
}

func modifyKeyForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("host").
				Title("Host").
				Placeholder("github.com-username"),
			huh.NewInput().
				Key("hostname").
				Title("HostName").
				Placeholder("github.com"),
			huh.NewInput().
				Key("user").
				Title("User").
				Placeholder("githubuser"),
			huh.NewText().
				Key("extras").
				Title("extras").
				Placeholder("IdentitiesOnly yes"),
		),
	)
}

func New() Modify {
	return Modify{
		form:         modifyKeyForm(),
		formFinished: false,
	}
}
