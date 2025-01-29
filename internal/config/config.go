package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_URL            string  `json:"db_url"`
	Current_user_name *string `json:"current_user_name,omitempty"`
}

func Read(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %v", err)
	}

  var config Config
  err = json.Unmarshal(b, &config)
  if err != nil {
    return nil, fmt.Errorf("Failed to unmarshal json: %v", err)
  }

	fmt.Println(config)

	return &Config{}, nil
}
