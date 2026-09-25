package interactive

import (
	"errors"
	"fmt"

	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/config"
	"github.com/sampaxyz/rolarte-cli/internal/prompt"
)

func Configure(
	application *app.App,
) error {
	current, err := application.LoadConfig()

	if err != nil &&
		!errors.Is(
			err,
			config.ErrSystemProjectsRootNotConfigured,
		) {
		return err
	}

	next, err := prompt.Configure(current)
	if err != nil {
		return err
	}

	if err := application.SaveConfig(next); err != nil {
		return err
	}

	path, err := application.ConfigPath()
	if err != nil {
		return err
	}

	fmt.Printf(
		"\nConfiguration saved to:\n%s\n",
		path,
	)

	return nil
}
