
package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#4F5D75")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 2)
)

// Header returns a styled header component
func Header(title string) string {
	return headerStyle.Render(title)
}
