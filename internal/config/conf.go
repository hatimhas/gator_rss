package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const cfgfileName = ".gatorconfig.json"

func Read() (Config, error) {
	cfgPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	// Open file using os.Open
	file, err := os.Open(cfgPath)
	if err != nil {
		return Config{}, err
	}

	// defer close file
	defer file.Close()

	var cfg Config
	// Decode the file(json) into cfg
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c *Config) Save() error {
	cfgPath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}

// SetUser sets the username and saves the config
func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return c.Save()
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cfgPath := filepath.Join(homeDir, cfgfileName)
	return cfgPath, nil
}

func (c Config) PrettyString() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "<error marshaling config>"
	}
	return string(data)
}
