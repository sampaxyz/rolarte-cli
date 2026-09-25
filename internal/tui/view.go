package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (model model) View() tea.View {
	content := model.render()

	view := tea.NewView(content)
	view.AltScreen = true

	return view
}

func (model model) render() string {
	if model.width == 0 {
		return ""
	}

	var content string

	switch model.screen {
	case screenLoading:
		content = model.renderLoading()

	case screenDashboard:
		content = model.renderDashboard()

	case screenProjectName:
		content = model.renderProjectName()

	case screenProjectSubprojects:
		content = model.renderProjectSubprojects()

	case screenProjectClipCount:
		content = model.renderProjectClipCount()

	case screenProjectPreview:
		content = model.renderProjectPreview()

	case screenClipProject:
		content = model.renderClipProject()

	case screenClipCount:
		content = model.renderClipCount()

	case screenClipPreview:
		content = model.renderClipPreview()

	case screenConfig:
		content = model.renderConfig()

	case screenSuccess:
		content = model.renderSuccess()

	case screenError:
		content = model.renderError()
	}

	return lipgloss.Place(
		model.width,
		model.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (model model) frame(
	breadcrumb string,
	content string,
	help string,
) string {
	var builder strings.Builder

	builder.WriteString(renderLogo())
	builder.WriteString("\n\n")

	builder.WriteString(
		breadcrumbStyle.Render(
			breadcrumb,
		),
	)

	if model.rootPath != "" {
		builder.WriteString("\n")

		builder.WriteString(
			dimStyle.Render(
				model.rootPath,
			),
		)
	}

	builder.WriteString("\n\n")
	builder.WriteString(content)
	builder.WriteString("\n\n")

	builder.WriteString(
		footerStyle.Render(help),
	)

	return builder.String()
}

func renderLogo() string {
	return brandStyle.Render(
		"◇━━━┫  R O L A R T E  ┣━━━━━━━━━━━━━━━━━━",
	)
}

func (model model) renderLoading() string {
	return model.frame(
		"ROLARTE",

		primaryPanelStyle.
			Width(60).
			Render(
				"Loading production workspace...",
			),

		"ctrl+c quit",
	)
}

func (model model) renderDashboard() string {
	menu := model.renderDashboardMenu()
	catalog := model.renderCatalog()

	left := panelStyle.
		Width(30).
		Render(menu)

	rightWidth := model.width - 48

	if rightWidth < 38 {
		rightWidth = 38
	}

	if rightWidth > 70 {
		rightWidth = 70
	}

	right := panelStyle.
		Width(rightWidth).
		Render(catalog)

	panels := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		"  ",
		right,
	)

	releaseMessage :=
		model.renderReleaseMessage()

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		panels,
		"",
		releaseMessage,
	)

	return model.frame(
		"ROLARTE  ›  SYSTEM PROJECTS",

		content,

		"↑/↓ navigate   enter select   q quit",
	)
}

func (model model) renderDashboardMenu() string {
	var builder strings.Builder

	builder.WriteString(
		titleStyle.Render("ACTIONS"),
	)

	builder.WriteString("\n\n")

	for index, item := range model.dashboardItems {

		label := "  " + item.label

		if index == model.dashboardCursor {
			label = "❯ " + item.label

			builder.WriteString(
				selectedStyle.Render(label),
			)
		} else {
			builder.WriteString(
				menuItemStyle.Render(label),
			)
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func (model model) renderCatalog() string {
	var builder strings.Builder

	builder.WriteString(
		titleStyle.Render(
			"RELEASE CATALOG",
		),
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

	maxVisible := model.height - 16

	if maxVisible < 3 {
		maxVisible = 3
	}

	visible := len(model.projects)

	if visible > maxVisible {
		visible = maxVisible
	}

	for index := 0; index < visible; index++ {
		fmt.Fprintf(
			&builder,
			"%02d  %s\n",
			index+1,
			model.projects[index],
		)
	}

	if visible < len(model.projects) {
		fmt.Fprintf(
			&builder,
			"\n%s",
			dimStyle.Render(
				fmt.Sprintf(
					"… %d more",
					len(model.projects)-visible,
				),
			),
		)
	}

	builder.WriteString("\n\n")

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
		return "Congrats. 1 production has made it through the pipeline."

	default:
		return fmt.Sprintf(
			"Congrats. %d productions have made it through the pipeline.",
			count,
		)
	}
}

func (model model) renderProjectName() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,

		titleStyle.Render(
			"NEW SYSTEM PROJECT",
		),

		"",

		"Give this production a name.",

		"",

		model.projectNameInput.View(),

		"",

		model.renderValidationError(),
	)

	return model.frame(
		"ROLARTE  ›  CREATE  ›  PROJECT  ›  NAME",

		primaryPanelStyle.
			Width(66).
			Render(content),

		"enter continue   esc back   ctrl+c quit",
	)
}

func (model model) renderProjectSubprojects() string {
	var builder strings.Builder

	builder.WriteString(
		titleStyle.Render(
			"SELECT SUBPROJECTS",
		),
	)

	builder.WriteString("\n\n")

	builder.WriteString(
		dimStyle.Render(
			"Space toggles a production module.",
		),
	)

	builder.WriteString("\n\n")

	for index, option := range projectOptions {

		check := "[ ]"

		if model.projectSelected[index] {
			check = "[x]"
		}

		line := fmt.Sprintf(
			"%s  %s",
			check,
			option.label,
		)

		if index ==
			model.projectSubprojectCursor {

			builder.WriteString(
				selectedStyle.Render(
					"❯ " + line,
				),
			)
		} else {
			builder.WriteString(
				"  " + line,
			)
		}

		builder.WriteString("\n")
	}

	if model.validationError != "" {
		builder.WriteString("\n")

		builder.WriteString(
			errorStyle.Render(
				model.validationError,
			),
		)
	}

	return model.frame(
		"ROLARTE  ›  CREATE  ›  PROJECT  ›  SUBPROJECTS",

		primaryPanelStyle.
			Width(66).
			Render(
				builder.String(),
			),

		"↑/↓ navigate   space toggle   enter continue   esc back",
	)
}

func (model model) renderProjectClipCount() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,

		titleStyle.Render(
			"INITIAL CLIPS",
		),

		"",

		"How many clips should this production start with?",

		"",

		model.projectClipCountInput.View(),

		"",

		model.renderValidationError(),
	)

	return model.frame(
		"ROLARTE  ›  CREATE  ›  PROJECT  ›  CLIPS",

		primaryPanelStyle.
			Width(66).
			Render(content),

		"enter preview   esc back",
	)
}

func (model model) renderProjectPreview() string {
	plan := model.projectPlan

	location := filepath.Join(
		plan.BasePath,
		plan.Plan.Root,
	)

	content := fmt.Sprintf(
		"%s\n\nProject   %s\nLocation  %s\n\n%s",
		titleStyle.Render(
			"GENERATION PLAN",
		),

		plan.Plan.Root,
		location,

		model.renderDirectories(
			plan.Plan.Directories,
			model.projectPreviewPage,
		),
	)

	return model.frame(
		"ROLARTE  ›  CREATE  ›  PROJECT  ›  PREVIEW",

		primaryPanelStyle.
			Width(78).
			Render(content),

		"←/→ pages   enter create   b/esc back",
	)
}

func (model model) renderClipProject() string {
	var builder strings.Builder

	builder.WriteString(
		titleStyle.Render(
			"SELECT SYSTEM PROJECT",
		),
	)

	builder.WriteString("\n\n")

	for index, projectName := range model.projects {

		line := fmt.Sprintf(
			"%02d  %s",
			index+1,
			projectName,
		)

		if index ==
			model.clipProjectCursor {

			builder.WriteString(
				selectedStyle.Render(
					"❯ " + line,
				),
			)
		} else {
			builder.WriteString(
				"  " + line,
			)
		}

		builder.WriteString("\n")
	}

	return model.frame(
		"ROLARTE  ›  CREATE  ›  CLIP  ›  PROJECT",

		primaryPanelStyle.
			Width(66).
			Render(
				builder.String(),
			),

		"↑/↓ navigate   enter continue   esc back",
	)
}

func (model model) renderClipCount() string {
	projectName :=
		model.projects[model.clipProjectCursor]

	content := lipgloss.JoinVertical(
		lipgloss.Left,

		titleStyle.Render(
			"ADDITIONAL CLIPS",
		),

		"",

		fmt.Sprintf(
			"Project: %s",
			projectName,
		),

		"",

		"How many new clips should be created?",

		"",

		model.clipCountInput.View(),

		"",

		model.renderValidationError(),
	)

	return model.frame(
		"ROLARTE  ›  CREATE  ›  CLIP  ›  COUNT",

		primaryPanelStyle.
			Width(66).
			Render(content),

		"enter preview   esc back",
	)
}

func (model model) renderClipPreview() string {
	plan := model.clipPlan

	content := fmt.Sprintf(
		"%s\n\nProject   %s\nLocation  %s\n\n%s",
		titleStyle.Render(
			"CLIP GENERATION PLAN",
		),

		plan.ProjectName,
		plan.ProjectPath,

		model.renderDirectories(
			plan.Plan.Directories,
			model.clipPreviewPage,
		),
	)

	return model.frame(
		"ROLARTE  ›  CREATE  ›  CLIP  ›  PREVIEW",

		primaryPanelStyle.
			Width(78).
			Render(content),

		"←/→ pages   enter create   b/esc back",
	)
}

func (model model) renderConfig() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,

		titleStyle.Render(
			"CONFIGURATION",
		),

		"",

		"System Projects Root",

		"",

		model.configRootInput.View(),

		"",

		dimStyle.Render(
			"All productions in the Release Catalog are discovered here.",
		),

		"",

		model.renderValidationError(),
	)

	help := "enter save   esc back"

	if model.rootPath == "" {
		help = "enter save   esc quit"
	}

	return model.frame(
		"ROLARTE  ›  CONFIGURATION",

		primaryPanelStyle.
			Width(76).
			Render(content),

		help,
	)
}

func (model model) renderSuccess() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,

		successStyle.Render(
			"✓ "+model.successTitle,
		),

		"",

		model.successMessage,

		"",

		dimStyle.Render(
			"The production workspace has been updated.",
		),
	)

	return model.frame(
		"ROLARTE  ›  COMPLETE",

		primaryPanelStyle.
			Width(70).
			Render(content),

		"enter return to dashboard",
	)
}

func (model model) renderError() string {
	message := "Unknown error."

	if model.currentError != nil {
		message =
			model.currentError.Error()
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,

		errorStyle.Render(
			"ERROR",
		),

		"",

		message,
	)

	return model.frame(
		"ROLARTE  ›  ERROR",

		primaryPanelStyle.
			Width(70).
			Render(content),

		"enter return to dashboard   q quit",
	)
}

func (model model) renderValidationError() string {
	if model.validationError == "" {
		return ""
	}

	return errorStyle.Render(
		model.validationError,
	)
}

func (model model) renderDirectories(
	directories []string,
	page int,
) string {
	var builder strings.Builder

	totalPages := pageCount(
		len(directories),
		previewPageSize,
	)

	currentPage := pageSlice(
		directories,
		page,
		previewPageSize,
	)

	builder.WriteString(
		titleStyle.Render(
			"DIRECTORIES",
		),
	)

	builder.WriteString("\n\n")

	for _, directory := range currentPage {
		fmt.Fprintf(
			&builder,
			"  %s/\n",
			directory,
		)
	}

	builder.WriteString("\n")

	builder.WriteString(
		dimStyle.Render(
			fmt.Sprintf(
				"Page %d/%d",
				page+1,
				totalPages,
			),
		),
	)

	return builder.String()
}
