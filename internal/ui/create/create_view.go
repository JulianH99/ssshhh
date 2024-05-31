package create

import (
	"fmt"

	"github.com/JulianH99/ssshhh/internal/config"
	"github.com/JulianH99/ssshhh/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Create struct {
	form         *huh.Form
	formFinished bool
	newKeyPath   string
}

func (m Create) GetKeyPath() string {
	return m.newKeyPath
}

func (m Create) Init() tea.Cmd {
	return m.form.Init()
}

func (m Create) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmd := m.form.Update(msg)

	m.form = newForm.(*huh.Form)

	if m.form.State == huh.StateCompleted && !m.formFinished {
		fmt.Println("this is it")
		m.formFinished = true
		return m, execSshKeygenCmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case execSshKeygen:
		var (
			path    = m.form.GetString("path")
			comment = m.form.GetString("comment")
			keyType = m.form.GetString("type")
		)

		createSshConfig := config.CreateSshKeyConfig(path, keyType, comment)
		createCommandExec := config.CreateSshKey(createSshConfig)

		return m, tea.ExecProcess(createCommandExec, func(err error) tea.Msg {
			fmt.Println("exec key has finished", err)
			return ui.KeyCreatedMsg{Key: createSshConfig.Path()}
		})
	}

	return m, cmd
}

func (m Create) View() string {
	return m.form.View()
}

func createKeyForm() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("path").
				Title("Name").
				Placeholder("Name of the SSH key"),
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

func New() Create {
	return Create{
		form:         createKeyForm(),
		formFinished: false,
	}
}
