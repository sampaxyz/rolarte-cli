package tui

import "charm.land/lipgloss/v2"

var (
	brandStyle = lipgloss.NewStyle().
			Bold(true)

	dimStyle = lipgloss.NewStyle().
			Faint(true)

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Reverse(true).
			Padding(0, 1)

	menuItemStyle = lipgloss.NewStyle().
			Padding(0, 1)

	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2)

	breadcrumbStyle = lipgloss.NewStyle().
			Bold(true)

	countStyle = lipgloss.NewStyle().
			Faint(true)
)
