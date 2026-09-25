package filesystem

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/sampaxyz/rolarte-cli/internal/project"
)

func TestExecutorCreatesProjectStructure(t *testing.T) {
	tempDir := t.TempDir()

	plan, err := project.BuildPlan(project.CreateRequest{
		Name: "Outta Bubblegum",
		Subprojects: []project.Subproject{
			project.SubprojectMicrofiction,
			project.SubprojectClips,
		},
		InitialClips: 1,
	})

	if err != nil {
		t.Fatalf("failed to build plan: %v", err)
	}

	executor := Executor{}

	rootPath, err := executor.Execute(tempDir, plan)
	if err != nil {
		t.Fatalf("failed to execute plan: %v", err)
	}

	expectedDirectories := []string{
		"1_MICROFICTION",
		"1_MICROFICTION/1.1_ASSETS",
		"4_CLIPS",
		"4_CLIPS/4.1_CLIP1",
		"4_CLIPS/4.1_CLIP1/4.1.4_EXPORTS",
	}

	for _, directory := range expectedDirectories {
		path := filepath.Join(rootPath, directory)

		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf(
				"expected directory %q to exist: %v",
				directory,
				err,
			)
		}

		if !info.IsDir() {
			t.Errorf(
				"expected %q to be a directory",
				directory,
			)
		}
	}
}

func TestExecutorRejectsExistingProject(t *testing.T) {
	tempDir := t.TempDir()

	plan, err := project.BuildPlan(project.CreateRequest{
		Name: "Outta Bubblegum",
	})

	if err != nil {
		t.Fatalf("failed to build plan: %v", err)
	}

	existingPath := filepath.Join(
		tempDir,
		"OUTTA_BUBBLEGUM",
	)

	if err := os.MkdirAll(existingPath, 0o755); err != nil {
		t.Fatalf("failed to prepare test: %v", err)
	}

	executor := Executor{}

	_, err = executor.Execute(tempDir, plan)

	if !errors.Is(err, ErrProjectAlreadyExists) {
		t.Fatalf(
			"expected ErrProjectAlreadyExists, got %v",
			err,
		)
	}
}
