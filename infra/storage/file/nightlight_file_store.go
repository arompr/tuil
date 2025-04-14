package file_storage

import (
	"fmt"
	"lighttui/domain/adjustable"
	"lighttui/domain/adjustable/nightlight"
	"os"
	"path/filepath"
	"strconv"
)

// FileNightLightStore handles reading, and storing the current hyprsunset temperature.
type FileNightLightStore struct {
	path string
}

// NewTemperatureStore initializes the store, reading the temperature or setting a default.
func NewHyprsunsetFileStore(filePath string) (*FileNightLightStore, error) {
	f := &FileNightLightStore{filePath}
	err := f.initFileStore()
	return f, err
}

// Reads the night light temperature file or initializes it with the min default value.
func (f *FileNightLightStore) initFileStore() error {
	dir := filepath.Dir(f.path)

	// Create all necessary directories if they don't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories for night light store: %w", err)
	}

	_, err := os.Stat(f.path)
	if os.IsNotExist(err) {
		return f.writeToFile(nightlight.MinTemperature)
	}

	if err != nil {
		return fmt.Errorf("error checking night light store file: %w", err)
	}

	return nil
}

func (f *FileNightLightStore) Fetch() (adjustable.IAdjustable, error) {
	temperature, err := f.readTemperature()

	if err != nil {
		return nil, fmt.Errorf("error fetching night light from file")
	}

	return nightlight.CreateNewNightLight(temperature), nil
}

// readTemperature reads the night light temperature from the file or sets it's value if file is empty.
func (f *FileNightLightStore) readTemperature() (int, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nightlight.MinTemperature, fmt.Errorf("failed to read file: %w", err)
	}

	// Handle empty file by setting default
	if len(data) == 0 {
		return nightlight.MinTemperature, f.writeToFile(nightlight.MinTemperature)
	}

	temperature, err := strconv.Atoi(string(data))
	if err != nil {
		return nightlight.MinTemperature, fmt.Errorf("invalid night light value in file: %w", err)
	}

	return temperature, nil
}

func (f *FileNightLightStore) Save(adjustable adjustable.IAdjustable) error {
	if err := f.writeToFile(adjustable.GetCurrentValue()); err != nil {
		return fmt.Errorf("failed to write temperature file: %w", err)
	}
	return nil
}

func (f *FileNightLightStore) writeToFile(value int) error {
	return os.WriteFile(f.path, []byte(strconv.Itoa(value)), 0644)
}
