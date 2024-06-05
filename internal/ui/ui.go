package ui

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// colors
const (
	cPrimary       = lipgloss.Color("#0099ff")
	cPrimaryOpaque = lipgloss.Color("#80ccff")
	cAccent        = lipgloss.Color("#cc99ff")
	cAccentOpaque  = lipgloss.Color("#e6ccff")
	cWhite         = lipgloss.Color("#ffffff")
	cBlack         = lipgloss.Color("#222222")
)

// styles
var (
	AppStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(cPrimary)

	FormTitleStyle = lipgloss.NewStyle().
			Foreground(cWhite).
			Background(cPrimary).
			MarginBottom(1)

	ListTitleStyle = lipgloss.NewStyle().
			Foreground(cWhite).
			Background(cPrimary).
			Bold(true)

	ListDescStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(cAccent)

	ListItemTitleStyle = lipgloss.NewStyle().
				Foreground(cPrimary)

	StatusMessageStyle = lipgloss.NewStyle().
				Background(lipgloss.AdaptiveColor{Light: "#0066ff", Dark: "#0099ff"}).
				Foreground(cBlack).
				Padding(0, 2).
				Render

	ListStatusBarStyle = lipgloss.NewStyle().
				Background(cAccent).
				Foreground(cBlack).
				MarginBottom(1)

	FormLabelStyle = lipgloss.NewStyle().
			Foreground(cPrimary)
)

func FormTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Title = t.Focused.Title.Foreground(cPrimary)
	t.Focused.Base = t.Focused.Base.BorderForeground(cAccent)
	t.Focused.Description = t.Focused.Description.Foreground(cAccent)
	t.Focused.Option = t.Focused.Option.Foreground(cAccentOpaque)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(cAccent)

	t.Blurred.Title = t.Blurred.Title.Foreground(cPrimaryOpaque)
	t.Blurred.Description = t.Blurred.Description.Foreground(cAccentOpaque)
	t.Blurred.Base = t.Blurred.Base.BorderForeground(cAccentOpaque)

	return t
}
