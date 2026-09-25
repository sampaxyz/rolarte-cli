package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListProjects(t *testing.T) {
	root := t.TempDir()

	expectedProjects := []string{
		"OUTTA_BUBBLEGUM",
		"THE_ONE_RING",
	}

	for _, projectName := range expectedProjects {
		if err := os.Mkdir(
			filepath.Join(root, projectName),
			0o755,
		); err != nil {
			t.Fatalf("failed to create project: %v", err)
		}
	}

	projects, err := ListProjects(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(projects) != 2 {
		t.Fatalf(
			"expected 2 projects, got %d",
			len(projects),
		)
	}
}

func TestNextClipIndex(t *testing.T) {
	projectPath := t.TempDir()

	clipsPath := filepath.Join(
		projectPath,
		"4_CLIPS",
	)

	existing := []string{
		"4.1_CLIP1",
		"4.2_CLIP2",
		"4.5_CLIP5",
	}

	for _, directory := range existing {
		if err := os.MkdirAll(
			filepath.Join(clipsPath, directory),
			0o755,
		); err != nil {
			t.Fatalf("failed to prepare test: %v", err)
		}
	}

	index, err := NextClipIndex(projectPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if index != 6 {
		t.Fatalf(
			"expected next clip index 6, got %d",
			index,
		)
	}
}

func TestNextClipIndexReturnsOneWithoutClipsFolder(
	t *testing.T,
) {
	projectPath := t.TempDir()

	index, err := NextClipIndex(projectPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if index != 1 {
		t.Fatalf(
			"expected next clip index 1, got %d",
			index,
		)
	}
}
