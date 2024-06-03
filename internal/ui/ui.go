package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// colors
const (
	C_PRIMARY = "#0099ff"
	C_ACCENT  = "#cc99ff"
)

// styles
var (
	AppStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(C_ACCENT))

	ListTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color(C_PRIMARY)).
			Bold(true)

	ListDescStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color(C_ACCENT))

	ListItemTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(C_PRIMARY))

	StatusMessageStyle = lipgloss.NewStyle().
				Background(lipgloss.AdaptiveColor{Light: "#0066ff", Dark: "#0099ff"}).
				Foreground(lipgloss.Color("#222222")).
				Padding(0, 2).
				Render

	ListStatusBarStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(C_ACCENT)).
				Foreground(lipgloss.Color("#222222")).
				MarginBottom(1)

	FormLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(C_PRIMARY))
)
