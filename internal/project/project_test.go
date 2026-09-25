package project

import (
	"errors"
	"testing"
)

func TestBuildPlanNormalizesProjectName(t *testing.T) {
	plan, err := BuildPlan(CreateRequest{
		Name: "Outta Bubblegum",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if plan.Root != "OUTTA_BUBBLEGUM" {
		t.Errorf(
			"expected OUTTA_BUBBLEGUM, got %s",
			plan.Root,
		)
	}
}

func TestBuildPlanCreatesMicrofictionDirectories(t *testing.T) {
	plan, err := BuildPlan(CreateRequest{
		Name: "Outta Bubblegum",
		Subprojects: []Subproject{
			SubprojectMicrofiction,
		},
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"1_MICROFICTION",
		"1_MICROFICTION/1.1_ASSETS",
		"1_MICROFICTION/1.2_AUDIO",
		"1_MICROFICTION/1.3_ANIMATION",
		"1_MICROFICTION/1.4_EXPORTS",
	}

	assertDirectories(t, plan.Directories, expected)
}

func TestBuildPlanCreatesInitialClip(t *testing.T) {
	plan, err := BuildPlan(CreateRequest{
		Name: "Outta Bubblegum",
		Subprojects: []Subproject{
			SubprojectClips,
		},
		InitialClips: 1,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"4_CLIPS",
		"4_CLIPS/4.1_CLIP1",
		"4_CLIPS/4.1_CLIP1/4.1.1_ASSETS",
		"4_CLIPS/4.1_CLIP1/4.1.2_AUDIO",
		"4_CLIPS/4.1_CLIP1/4.1.3_ANIMATION",
		"4_CLIPS/4.1_CLIP1/4.1.4_EXPORTS",
	}

	assertDirectories(t, plan.Directories, expected)
}

func TestBuildPlanRejectsEmptyName(t *testing.T) {
	_, err := BuildPlan(CreateRequest{
		Name: "   ",
	})

	if !errors.Is(err, ErrEmptyProjectName) {
		t.Fatalf(
			"expected ErrEmptyProjectName, got %v",
			err,
		)
	}
}

func assertDirectories(
	t *testing.T,
	actual []string,
	expected []string,
) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf(
			"expected %d directories, got %d",
			len(expected),
			len(actual),
		)
	}

	for index := range expected {
		if actual[index] != expected[index] {
			t.Errorf(
				"directory %d: expected %q, got %q",
				index,
				expected[index],
				actual[index],
			)
		}
	}
}
