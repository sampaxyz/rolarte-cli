package prompt

import (
	"charm.land/huh/v2"

	"github.com/sampaxyz/rolarte-cli/internal/config"
)

func Configure(
	current config.Config,
) (config.Config, error) {
	systemProjectsRoot := current.SystemProjectsRoot

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("System Projects Root").
				Description(
					"Directory where Rolarte System Projects are stored.",
				).
				Value(&systemProjectsRoot).
				Validate(validateRequired),
		),
	)

	if err := form.Run(); err != nil {
		return config.Config{}, err
	}

	return config.Config{
		SystemProjectsRoot: systemProjectsRoot,
	}, nil
}
