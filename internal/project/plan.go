package project

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrEmptyProjectName = errors.New("project name cannot be empty")
	ErrInvalidClipCount = errors.New("initial clips cannot be negative")
)

type Plan struct {
	Root        string
	Directories []string
}

func BuildPlan(request CreateRequest) (Plan, error) {
	root := normalizeProjectName(request.Name)

	if root == "" {
		return Plan{}, ErrEmptyProjectName
	}

	if request.InitialClips < 0 {
		return Plan{}, ErrInvalidClipCount
	}

	directories := make([]string, 0)

	for _, subproject := range request.Subprojects {
		switch subproject {
		case SubprojectMicrofiction:
			directories = append(
				directories,
				microfictionDirectories()...,
			)

		case SubprojectSystemOverview:
			directories = append(
				directories,
				systemOverviewDirectories()...,
			)

		case SubprojectPodcastGame:
			directories = append(
				directories,
				gameSessionDirectories()...,
			)

		case SubprojectPodcastRolafter:
			directories = append(
				directories,
				rolafterDirectories()...,
			)

		case SubprojectClips:
			directories = append(
				directories,
				clipDirectories(request.InitialClips)...,
			)
		}
	}

	return Plan{
		Root:        root,
		Directories: uniqueDirectories(directories),
	}, nil
}

func normalizeProjectName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)

	separator := regexp.MustCompile(`[^A-Z0-9]+`)
	name = separator.ReplaceAllString(name, "_")

	return strings.Trim(name, "_")
}

func microfictionDirectories() []string {
	return []string{
		"1_MICROFICTION",
		"1_MICROFICTION/1.1_ASSETS",
		"1_MICROFICTION/1.2_AUDIO",
		"1_MICROFICTION/1.3_ANIMATION",
		"1_MICROFICTION/1.4_EXPORTS",
	}
}

func systemOverviewDirectories() []string {
	return []string{
		"2_SYSTEM_OVERVIEW",
		"2_SYSTEM_OVERVIEW/2.1_ASSETS",
		"2_SYSTEM_OVERVIEW/2.2_AUDIO",
		"2_SYSTEM_OVERVIEW/2.3_REEL",
		"2_SYSTEM_OVERVIEW/2.4_EXPORTS",
	}
}

func gameSessionDirectories() []string {
	return []string{
		"3_PODCAST",
		"3_PODCAST/3.1_GAME_SESSION",
		"3_PODCAST/3.1_GAME_SESSION/3.1.1_AUDIO",
		"3_PODCAST/3.1_GAME_SESSION/3.1.2_VIDEO",
		"3_PODCAST/3.1_GAME_SESSION/3.1.3_EXPORTS",
	}
}

func rolafterDirectories() []string {
	return []string{
		"3_PODCAST",
		"3_PODCAST/3.2_ROLAFTER",
		"3_PODCAST/3.2_ROLAFTER/3.2.1_AUDIO",
		"3_PODCAST/3.2_ROLAFTER/3.2.2_VIDEO",
		"3_PODCAST/3.2_ROLAFTER/3.2.3_EXPORTS",
	}
}

func clipDirectories(count int) []string {
	if count == 0 {
		return []string{"4_CLIPS"}
	}

	directories := []string{
		"4_CLIPS",
	}

	for index := 1; index <= count; index++ {
		root := fmt.Sprintf(
			"4_CLIPS/4.%d_CLIP%d",
			index,
			index,
		)

		directories = append(
			directories,
			root,
			fmt.Sprintf("%s/4.%d.1_ASSETS", root, index),
			fmt.Sprintf("%s/4.%d.2_AUDIO", root, index),
			fmt.Sprintf("%s/4.%d.3_ANIMATION", root, index),
			fmt.Sprintf("%s/4.%d.4_EXPORTS", root, index),
		)
	}

	return directories
}

func uniqueDirectories(directories []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(directories))

	for _, directory := range directories {
		if _, exists := seen[directory]; exists {
			continue
		}

		seen[directory] = struct{}{}
		result = append(result, directory)
	}

	return result
}
