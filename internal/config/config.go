package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const cfgFilename = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadConfig() (*Config, error) {
	cfg := &Config{}

	cfgFile, err := os.Open(getCfgFilePath())
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	err = json.NewDecoder(cfgFile).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	return writeConfig(c)
}

func writeConfig(c *Config) error {
	cfgFile, err := os.Create(getCfgFilePath())
	if err != nil {
		return err
	}
	defer cfgFile.Close()
	
	return json.NewEncoder(cfgFile).Encode(c)
}

func getCfgFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, cfgFilename)
}
