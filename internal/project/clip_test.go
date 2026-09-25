package project

import (
	"errors"
	"testing"
)

func TestBuildClipPlan(t *testing.T) {
	plan, err := BuildClipPlan(3, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if plan.StartIndex != 3 {
		t.Fatalf(
			"expected start index 3, got %d",
			plan.StartIndex,
		)
	}

	expected := []string{
		"4_CLIPS/4.3_CLIP3",
		"4_CLIPS/4.3_CLIP3/4.3.1_ASSETS",
		"4_CLIPS/4.3_CLIP3/4.3.2_AUDIO",
		"4_CLIPS/4.3_CLIP3/4.3.3_ANIMATION",
		"4_CLIPS/4.3_CLIP3/4.3.4_EXPORTS",

		"4_CLIPS/4.4_CLIP4",
		"4_CLIPS/4.4_CLIP4/4.4.1_ASSETS",
		"4_CLIPS/4.4_CLIP4/4.4.2_AUDIO",
		"4_CLIPS/4.4_CLIP4/4.4.3_ANIMATION",
		"4_CLIPS/4.4_CLIP4/4.4.4_EXPORTS",
	}

	assertDirectories(
		t,
		plan.Directories,
		expected,
	)
}

func TestBuildClipPlanRejectsZeroCount(t *testing.T) {
	_, err := BuildClipPlan(1, 0)

	if !errors.Is(err, ErrInvalidAdditionalClipCount) {
		t.Fatalf(
			"expected ErrInvalidAdditionalClipCount, got %v",
			err,
		)
	}
}
