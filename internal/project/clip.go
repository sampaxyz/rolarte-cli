package project

import (
	"errors"
	"fmt"
)

var ErrInvalidAdditionalClipCount = errors.New(
	"additional clip count must be greater than zero",
)

type ClipPlan struct {
	StartIndex  int
	Directories []string
}

func BuildClipPlan(
	startIndex int,
	count int,
) (ClipPlan, error) {
	if startIndex < 1 {
		startIndex = 1
	}

	if count < 1 {
		return ClipPlan{}, ErrInvalidAdditionalClipCount
	}

	directories := make([]string, 0, count*5)

	for index := startIndex; index < startIndex+count; index++ {
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

	return ClipPlan{
		StartIndex:  startIndex,
		Directories: directories,
	}, nil
}
