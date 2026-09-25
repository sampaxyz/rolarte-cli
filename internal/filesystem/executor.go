package filesystem

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/sampaxyz/rolarte-cli/internal/project"
)

var ErrProjectAlreadyExists = errors.New("project already exists")

type Executor struct{}

func New() Executor {
	return Executor{}
}

func (Executor) Execute(
	basePath string,
	plan project.Plan,
) (string, error) {
	rootPath := filepath.Join(basePath, plan.Root)

	if _, err := os.Stat(rootPath); err == nil {
		return "", ErrProjectAlreadyExists
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	if err := os.MkdirAll(rootPath, 0o755); err != nil {
		return "", err
	}

	for _, directory := range plan.Directories {
		path := filepath.Join(rootPath, directory)

		if err := os.MkdirAll(path, 0o755); err != nil {
			return "", err
		}
	}

	return rootPath, nil
}
