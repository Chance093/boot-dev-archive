package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_URL            string `json:"db_url"`
	Current_user_name string `json:"current_user_name,omitempty"`
}

// getConfigFilePath is a helper func which returns the absolute path
// of gatorconfig json file.
func getConfigFilePath() (string, error) {
  const configFile = "/.gatorconfig.json"

  path, err := os.UserHomeDir()
  if err != nil {
    return "", fmt.Errorf("Failed to get config path: %v", err)
  }

  return path + configFile, nil
}

// Read will find the gator config file and read to a config struct
// which is returned to caller.
func Read() (*Config, error) {
  path, err := getConfigFilePath()
  if err != nil {
		return nil, err
  }

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal json: %v", err)
	}

	return &cfg, nil
}

// SetUser will set the username in the gator config file.
func (cfg *Config) SetUser(username string) error {
	cfg.Current_user_name = username
  return write(cfg)
}

// write is a helper function that overwrites the gator config file
// with the current config struct.
func write(cfg *Config) error {
  path, err := getConfigFilePath()
  if err != nil {
    return err
  }

	b, err := json.Marshal(*cfg)
	if err != nil {
		return fmt.Errorf("Failed to marshal json: %v", err)
	}

	err = os.WriteFile(path, b, 0666)
	if err != nil {
		return fmt.Errorf("Failed to write config file: %v", err)
	}

	return nil
}
