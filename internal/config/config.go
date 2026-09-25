package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var ErrSystemProjectsRootNotConfigured = errors.New(
	"system projects root is not configured",
)

type Config struct {
	SystemProjectsRoot string `json:"system_projects_root"`
}

func Path() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(
		configDir,
		"rolarte",
		"config.json",
	), nil
}

func Load() (Config, error) {
	path, err := Path()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, ErrSystemProjectsRootNotConfigured
		}

		return Config{}, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	if cfg.SystemProjectsRoot == "" {
		return Config{}, ErrSystemProjectsRootNotConfigured
	}

	return cfg, nil
}

func Save(cfg Config) error {
	path, err := Path()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(
		filepath.Dir(path),
		0o755,
	); err != nil {
		return err
	}

	data, err := json.MarshalIndent(
		cfg,
		"",
		"  ",
	)
	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		data,
		0o644,
	)
}
