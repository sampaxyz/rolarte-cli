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

	primaryPanelStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder()).
				Padding(1, 2)

	breadcrumbStyle = lipgloss.NewStyle().
			Bold(true)

	titleStyle = lipgloss.NewStyle().
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Bold(true)

	countStyle = lipgloss.NewStyle().
			Faint(true)

	checkStyle = lipgloss.NewStyle().
			Bold(true)

	footerStyle = lipgloss.NewStyle().
			Faint(true)
)
