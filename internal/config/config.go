package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	URL		string		`json:"db_url"`
	UserName	string		`json:"current_user_name"`
}

// helper function to get the full path 

func getConfigFilePath() (string,error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configpath := configFileName
	
	return filepath.Join(home, configpath), nil
}

func Read() (Config,error) {
	
	fullPath,err := getConfigFilePath()		
	if err != nil {
		return Config{}, err
	}

	val, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	gatorfile := Config{}
	err = json.Unmarshal(val, &gatorfile)
	if err != nil {
		return Config{},err
	}

	return gatorfile, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	updatedGatorFile, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, updatedGatorFile,0600)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.UserName = username
	err :=	write(*c)
	if err != nil {
		return err
	}
	return nil
}
