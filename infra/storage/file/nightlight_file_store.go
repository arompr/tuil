package file_storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"lighttui/domain/adjustable"
	"lighttui/domain/adjustable/nightlight"
)

type NightLightState struct {
	Enabled     bool `json:"enabled"`
	Temperature int  `json:"temperature"`
}

type FileNightLightStore struct {
	path string
}

// NewHyprsunsetFileStore ensures the file exists with default state.
func NewHyprsunsetFileStore(filePath string) (*FileNightLightStore, error) {
	f := &FileNightLightStore{filePath}
	if err := f.initFileStore(); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FileNightLightStore) initFileStore() error {
	dir := filepath.Dir(f.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create directories for night light store: %w", err)
	}

	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		defaultState := NightLightState{
			Enabled:     true,
			Temperature: nightlight.MinTemperature,
		}
		return f.writeState(defaultState)
	}
	return nil
}

func (f *FileNightLightStore) Fetch() (adjustable.IAdjustable, error) {
	state, err := f.readState()
	if err != nil {
		return nil, fmt.Errorf("error fetching night light: %w", err)
	}

	return toNightLight(state), nil
}

func toNightLight(state NightLightState) adjustable.IAdjustable {
	temp := state.Temperature
	isEnabled := state.Enabled
	return nightlight.CreateNewNightLight(temp, nightlight.WithEnabled(isEnabled))
}

func (f *FileNightLightStore) Save(adjustable adjustable.IAdjustable) error {
	state, err := f.readState()
	if err != nil {
		return err
	}

	state.Temperature = adjustable.GetCurrentValue()
	return f.writeState(state)
}

func (f *FileNightLightStore) readState() (NightLightState, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return NightLightState{}, fmt.Errorf("failed to read file: %w", err)
	}

	var state NightLightState
	if err := json.Unmarshal(data, &state); err != nil {
		// Reset if corrupted
		state = NightLightState{Enabled: true, Temperature: nightlight.MinTemperature}
		_ = f.writeState(state)
	}
	return state, nil
}

func (f *FileNightLightStore) writeState(state NightLightState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize night light state: %w", err)
	}
	return os.WriteFile(f.path, data, 0o644)
}
