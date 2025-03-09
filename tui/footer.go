package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	footerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#4F5D75")).
		Padding(0, 1).
		Width(80)
)

// Footer returns a styled footer with help text
func Footer(helpText string) string {
	return footerStyle.Render(helpText)
}
