package infra

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const storeFilePath = "~/.local/state/tuil_store"

func ReadOrInitTemperature() (string, error) {
	path := getTemperatureFilePath()

	// Check if file exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Create file with default value
		if err := os.WriteFile(path, []byte("6500"), 0644); err != nil {
			return "", fmt.Errorf("failed to create temp file: %w", err)
		}
		return "6500", nil
	} else if err != nil {
		return "", fmt.Errorf("error checking file: %w", err)
	}

	// Read file content
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read temp file: %w", err)
	}

	if len(data) == 0 {
		// Write default if empty
		if err := os.WriteFile(path, []byte("6500"), 0644); err != nil {
			return "", fmt.Errorf("failed to write default value: %w", err)
		}
		return "6500", nil
	}

	return string(data), nil
}

func Save(value int) {
	os.WriteFile(getTemperatureFilePath(), []byte(strconv.Itoa(value)), 0644)
}

func getTemperatureFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot determine home directory")
	}
	return filepath.Join(homeDir, ".local/state/hyprsunset_temp")
}
