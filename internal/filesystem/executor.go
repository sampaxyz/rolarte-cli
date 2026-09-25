package filesystem

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sampaxyz/rolarte-cli/internal/project"
)

var ErrProjectAlreadyExists = errors.New("project already exists")

type Executor struct{}

func NewExecutor() Executor {
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

func (Executor) ExecuteDirectories(
	rootPath string,
	directories []string,
) error {
	for _, directory := range directories {
		path := filepath.Join(
			rootPath,
			directory,
		)

		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf(
				"directory already exists: %s",
				path,
			)
		} else if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	for _, directory := range directories {
		path := filepath.Join(
			rootPath,
			directory,
		)

		if err := os.MkdirAll(
			path,
			0o755,
		); err != nil {
			return err
		}
	}

	return nil
}
