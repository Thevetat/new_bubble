package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Couldn't get user home directory: %s", err)
	}
	configPath := filepath.Join(home, ".config", "unfukurmac", "unfkurmacrc.json")
	return configPath, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CheckConfig() bool {
	configPath, err := getConfigPath()
	if err != nil {
		return false
	}
	configExists := FileExists(configPath)
	return configExists
}

type Config struct{}

func CreateConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	configDir := filepath.Dir(configPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("Couldn't create directory path: %s", err)
		}
	}

	defaultConfig := Config{}

	configFile, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("Couldn't create config file: %s", err)
	}

	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(&defaultConfig); err != nil {
		return fmt.Errorf("Couldn't encode config file: %s", err)
	}

	return nil
}
