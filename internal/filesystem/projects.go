package filesystem

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var clipDirectoryPattern = regexp.MustCompile(
	`^4\.(\d+)_CLIP\d+$`,
)

func ListProjects(basePath string) ([]string, error) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	projects := make([]string, 0)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		projects = append(
			projects,
			entry.Name(),
		)
	}

	return projects, nil
}

func NextClipIndex(projectPath string) (int, error) {
	clipsPath := filepath.Join(
		projectPath,
		"4_CLIPS",
	)

	entries, err := os.ReadDir(clipsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return 1, nil
		}

		return 0, err
	}

	maxIndex := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		matches := clipDirectoryPattern.FindStringSubmatch(
			entry.Name(),
		)

		if len(matches) != 2 {
			continue
		}

		index, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}

		if index > maxIndex {
			maxIndex = index
		}
	}

	return maxIndex + 1, nil
}
