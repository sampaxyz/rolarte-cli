package cli

import (
	"fmt"
	"strconv"
	"strings"

	"charm.land/huh/v2"
)

func validateRequired(value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("value cannot be empty")
	}

	return nil
}

func validatePositiveInt(value string) error {
	number, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return fmt.Errorf("must be a number")
	}

	if number < 1 {
		return fmt.Errorf("must be greater than zero")
	}

	return nil
}

func parseInt(value string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(value))
}

func confirmDirectories(
	title string,
	description string,
	directories []string,
) (bool, error) {
	var builder strings.Builder

	builder.WriteString(description)
	builder.WriteString("\n\nDirectories:\n")

	for _, directory := range directories {
		fmt.Fprintf(
			&builder,
			"  • %s\n",
			directory,
		)
	}

	var confirmed bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Description(builder.String()).
				Affirmative("Create").
				Negative("Cancel").
				Value(&confirmed),
		),
	)

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirmed, nil
}
