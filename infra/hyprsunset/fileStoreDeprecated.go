package hyprsunset

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// TemperatureStoreDeprecated handles reading, writing, and storing the current temperature.
type TemperatureStoreDeprecated struct {
	filePath           string
	currentTemperature int
}

// NewTemperatureStore initializes the store, reading the temperature or setting a default.
func NewTemperatureStore() (*TemperatureStoreDeprecated, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine home directory: %w", err)
	}

	store := &TemperatureStoreDeprecated{
		filePath: filepath.Join(homeDir, ".local/state/hyprsunset_temp"),
	}

	if err := store.initTemperature(); err != nil {
		return nil, err
	}

	return store, nil
}

// initTemperature reads the file or initializes it with a default value.
func (ts *TemperatureStoreDeprecated) initTemperature() error {
	_, err := os.Stat(ts.filePath)
	if os.IsNotExist(err) {
		return ts.writeTemperature(6500) // Set default value if file doesn't exist
	} else if err != nil {
		return fmt.Errorf("error checking temperature file: %w", err)
	}

	return ts.readTemperature()
}

// readTemperature reads the temperature from the file and updates currentTemperature.
func (ts *TemperatureStoreDeprecated) readTemperature() error {
	data, err := os.ReadFile(ts.filePath)
	if err != nil {
		return fmt.Errorf("failed to read temperature file: %w", err)
	}

	if len(data) == 0 {
		return ts.writeTemperature(6500) // Handle empty file by setting default
	}

	value, err := strconv.Atoi(string(data))
	if err != nil {
		return fmt.Errorf("invalid temperature value in file: %w", err)
	}

	ts.currentTemperature = value
	return nil
}

// writeTemperature updates the file and currentTemperature field.
func (ts *TemperatureStoreDeprecated) writeTemperature(value int) error {
	err := os.WriteFile(ts.filePath, []byte(strconv.Itoa(value)), 0644)
	if err != nil {
		return fmt.Errorf("failed to write temperature file: %w", err)
	}
	ts.currentTemperature = value
	return nil
}

// GetTemperature returns the current temperature.
func (ts *TemperatureStoreDeprecated) GetTemperature() int {
	return ts.currentTemperature
}

// Save updates the temperature and writes it to the file.
func (ts *TemperatureStoreDeprecated) Save(value int) error {
	return ts.writeTemperature(value)
}
