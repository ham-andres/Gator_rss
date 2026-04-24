package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DURL		string		`json:"db_url"`
	CurrentUserName	string		`json:"current_user_name"`
}

// helper function to get the full path 

func getConfigFilePath() (string,error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath :=  filepath.Join(home, configFileName)
	return fullPath, nil
}

func Read() (Config,error) {
	
	fullPath,err := getConfigFilePath()		
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	gatorfile := Config{}
	err = decoder.Decode(&gatorfile)
	if err != nil {
		return Config{},err
	}

	return gatorfile, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	updatedGatorFile, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, updatedGatorFile,0600)
	
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName	
	return write(*cfg)
}
