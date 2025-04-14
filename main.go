package main

import (
	"fmt"
	"lighttui/application/usecase"
	"lighttui/domain/adjustable/brightness"
	"lighttui/infra/brightnessctl"
	"lighttui/infra/hyprsunset"
	cached_storage "lighttui/infra/storage/cache"
	file_storage "lighttui/infra/storage/file"
	in_memory_storage "lighttui/infra/storage/in_memory"
	"lighttui/ui"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tui, err := initTuil()
	if err != nil {
		fmt.Println("Failed to initialize tuil: ", err)
		os.Exit(1)
	}

	if _, err := tui.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

func initTuil() (*tea.Program, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine home directory: %w", err)
	}

	// Initialise dependencies
	fileNightLightStore, err := file_storage.NewHyprsunsetFileStore(filepath.Join(homeDir, ".local/state/tuil/nightlight_temp"))
	if err != nil {
		return nil, err
	}

	inMemoryNightLightStore := in_memory_storage.NewInMemoryNightLightStore()
	cachedNightLightStore := cached_storage.NewCachedNightLightStore(inMemoryNightLightStore, fileNightLightStore)

	hyprsunsetAdapter, err := hyprsunset.NewNighLightAdapter(inMemoryNightLightStore)
	if err != nil {
		return nil, err
	}

	brightnessctlAdapter := brightnessctl.NewBrightnessCtlAdapter()
	currentBrightness, err := brightnessctlAdapter.GetCurrentBrightnessValue()
	if err != nil {
		return nil, err
	}

	maxBrightness, err := brightnessctlAdapter.GetMaxBrightnessValue()
	if err != nil {
		return nil, err
	}

	inMemoryBrightessStore := in_memory_storage.NewInMemoryBrightnessStore()
	inMemoryBrightessStore.Save(brightness.CreateNewBrightness(currentBrightness, maxBrightness))

	increaseNightLightUseCase := usecase.NewIncreaseUseCase(cachedNightLightStore, hyprsunsetAdapter)
	decreaseNightLightUseCase := usecase.NewDecreaseUseCase(cachedNightLightStore, hyprsunsetAdapter)
	getNightLightPercentageUseCase := usecase.NewGetPercentageUseCase(cachedNightLightStore)
	increaseBrightnessUseCase := usecase.NewIncreaseUseCase(inMemoryBrightessStore, brightnessctlAdapter)
	decreaseBrightnessUseCase := usecase.NewDecreaseUseCase(inMemoryBrightessStore, brightnessctlAdapter)
	getBrightnessPercentageUseCase := usecase.NewGetPercentageUseCase(inMemoryBrightessStore)
	persistNightLightUseCase := usecase.NewPersistUseCase(cachedNightLightStore, fileNightLightStore)

	return ui.NewTUI(
		increaseNightLightUseCase,
		decreaseNightLightUseCase,
		getNightLightPercentageUseCase,
		increaseBrightnessUseCase,
		decreaseBrightnessUseCase,
		getBrightnessPercentageUseCase,
		persistNightLightUseCase,
	), nil
}
