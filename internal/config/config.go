package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)


type Config struct {
	URL		string		`json:"db_url"`
	UserName	string		`json:"current_user_name"`
}

func Read() (Config,error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{},err
	}
	relativePath := ".gatorconfig.json"

	fullPath := filepath.Join(home, relativePath)
		
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
