package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (model model) View() tea.View {
	content := model.renderDashboard()

	view := tea.NewView(content)
	view.AltScreen = true

	return view
}

func (model model) renderDashboard() string {
	if model.width == 0 {
		return ""
	}

	logo := renderLogo()
	breadcrumb := renderBreadcrumb(model.rootPath)

	menu := model.renderMenu()
	catalog := model.renderCatalog()

	left := panelStyle.
		Width(30).
		Render(menu)

	rightWidth := max(38, model.width-42)

	right := panelStyle.
		Width(rightWidth).
		Render(catalog)

	panels := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		"  ",
		right,
	)

	footer := dimStyle.Render(
		"↑/↓ navigate   enter select   q quit",
	)

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		logo,
		"",
		breadcrumb,
		"",
		panels,
		"",
		model.renderReleaseMessage(),
		"",
		footer,
	)

	return lipgloss.Place(
		model.width,
		model.height,
		lipgloss.Center,
		lipgloss.Center,
		body,
	)
}

func renderLogo() string {
	return brandStyle.Render(
		"◇━━━┫  R O L A R T E  ┣━━━━━━━━━",
	)
}

func renderBreadcrumb(rootPath string) string {
	return breadcrumbStyle.Render(
		fmt.Sprintf(
			"~/ROLARTE  ›  %s",
			filepath.Base(rootPath),
		),
	)
}

func (model model) renderMenu() string {
	var builder strings.Builder

	builder.WriteString(
		brandStyle.Render("ACTIONS"),
	)
	builder.WriteString("\n\n")

	for index, item := range model.items {
		if index == model.cursor {
			builder.WriteString(
				selectedStyle.Render(
					"❯ " + item.label,
				),
			)
		} else {
			builder.WriteString(
				menuItemStyle.Render(
					"  " + item.label,
				),
			)
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func (model model) renderCatalog() string {
	var builder strings.Builder

	builder.WriteString(
		brandStyle.Render("RELEASE CATALOG"),
	)
	builder.WriteString("\n\n")

	if len(model.projects) == 0 {
		builder.WriteString(
			dimStyle.Render(
				"No systems released yet.",
			),
		)

		return builder.String()
	}

	for index, projectName := range model.projects {
		fmt.Fprintf(
			&builder,
			"%02d  %s\n",
			index+1,
			projectName,
		)
	}

	builder.WriteString("\n")

	builder.WriteString(
		countStyle.Render(
			fmt.Sprintf(
				"%d systems released",
				len(model.projects),
			),
		),
	)

	return builder.String()
}

func (model model) renderReleaseMessage() string {
	count := len(model.projects)

	switch count {
	case 0:
		return dimStyle.Render(
			"Your release catalog is waiting for its first production.",
		)

	case 1:
		return fmt.Sprintf(
			"Congrats. %d production has made it through the pipeline.",
			count,
		)

	default:
		return fmt.Sprintf(
			"Congrats. %d productions have made it through the pipeline.",
			count,
		)
	}
}
