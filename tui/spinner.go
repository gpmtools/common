package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Spinner is a reusable spinner component
type Spinner struct {
	spinner spinner.Model
	label   string
	style   lipgloss.Style
}

// NewSpinner creates a new spinner with a label
func NewSpinner(label string) Spinner {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	
	return Spinner{
		spinner: s,
		label:   label,
		style:   lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	}
}

// Init initializes the spinner
func (s Spinner) Init() tea.Cmd {
	return s.spinner.Tick
}

// Update handles messages for the spinner
func (s Spinner) Update(msg tea.Msg) (Spinner, tea.Cmd) {
	var cmd tea.Cmd
	s.spinner, cmd = s.spinner.Update(msg)
	return s, cmd
}

// View renders the spinner
func (s Spinner) View() string {
	return s.spinner.View() + " " + s.style.Render(s.label)
}
