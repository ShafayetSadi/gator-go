package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func (c *Config) Read() error {
	config_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("get config file path: %w", err)
	}

	data, err := os.ReadFile(config_path)
	if err != nil {
		return fmt.Errorf("read config file %q: %w", config_path, err)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("unmarshal config file %q: %w", config_path, err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Failed to get HOME dir")
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	config_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("get config file path: %w", err)
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	err = os.WriteFile(config_path, data, 0644)
	if err != nil {
		return fmt.Errorf("write config file %q: %w", config_path, err)
	}

	return nil
}
