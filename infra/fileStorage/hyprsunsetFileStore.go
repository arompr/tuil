package file_storage

import (
	"fmt"
	"lighttui/domain/nightlight"
	"os"
	"strconv"
)

// FileNightLightStore handles reading, and storing the current hyprsunset temperature.
type FileNightLightStore struct {
	filePath string
}

const maxTemperature = 1500
const minTemperature = 6500

// NewTemperatureStore initializes the store, reading the temperature or setting a default.
func NewHyprsunsetFileStore(filePath string) (*FileNightLightStore, error) {
	f := &FileNightLightStore{filePath}
	err := f.initFileStore()
	return f, err
}

// initFileStore reads the night light temperature file or initializes it with a default value.
func (f *FileNightLightStore) initFileStore() error {
	_, err := os.Stat(f.filePath)
	if os.IsNotExist(err) {
		return f.writeToFile(minTemperature)
	}

	if err != nil {
		return fmt.Errorf("error checking night light store file: %w", err)
	}

	return nil
}

func (f *FileNightLightStore) FetchNightLight() (*nightlight.NightLight, error) {
	temperature, err := f.readTemperature()

	if err != nil {
		return nil, fmt.Errorf("error fetching night light from file")
	}

	return nightlight.CreateNewNightLight(temperature, maxTemperature, minTemperature), nil
}

// readTemperature reads the night light temperature from the file or sets it's value if file is empty.
func (f *FileNightLightStore) readTemperature() (int, error) {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return minTemperature, fmt.Errorf("failed to read file: %w", err)
	}

	// Handle empty file by setting default
	if len(data) == 0 {
		return minTemperature, f.writeToFile(minTemperature)
	}

	temperature, err := strconv.Atoi(string(data))
	if err != nil {
		return minTemperature, fmt.Errorf("invalid night light value in file: %w", err)
	}

	return temperature, nil
}

func (f *FileNightLightStore) Save(nightlight *nightlight.NightLight) error {
	if err := f.writeToFile(nightlight.GetCurrentTemperature()); err != nil {
		return fmt.Errorf("failed to write temperature file: %w", err)
	}
	return nil
}

func (f *FileNightLightStore) writeToFile(value int) error {
	return os.WriteFile(f.filePath, []byte(strconv.Itoa(value)), 0644)
}
