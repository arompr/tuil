package main

import (
	"fmt"
	"os"
	"path/filepath"

	"lighttui/application/startup"
	"lighttui/application/usecase"
	"lighttui/domain/adjustable/brightness"
	"lighttui/infra/brightnessctl"
	"lighttui/infra/hyprsunset"
	cached_storage "lighttui/infra/storage/cache"
	file_storage "lighttui/infra/storage/file"
	in_memory_storage "lighttui/infra/storage/in_memory"
	"lighttui/ui"

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

	inMemoryNightLightStore := in_memory_storage.NewInMemoryNightlightStore()
	cachedNightLightStore := cached_storage.NewCachedNightlightStore(inMemoryNightLightStore, fileNightLightStore)

	hyprsunsetAdapter := hyprsunset.NewHyprsunsetAdapter()

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

	cachePersister := cached_storage.NewCachePersister(cachedNightLightStore)

	startNightlightServices := startup.NewStartNightlightServices(hyprsunsetAdapter, cachedNightLightStore)
	increaseNightLightUseCase := usecase.NewIncreaseNightlightUseCase(cachedNightLightStore, hyprsunsetAdapter)
	decreaseNightLightUseCase := usecase.NewDecreaseNightlightUseCase(cachedNightLightStore, hyprsunsetAdapter)
	getNightLightPercentageUseCase := usecase.NewGetNightlightPercentageUseCase(cachedNightLightStore)
	increaseBrightnessUseCase := usecase.NewIncreaseUseCase(inMemoryBrightessStore, brightnessctlAdapter)
	decreaseBrightnessUseCase := usecase.NewDecreaseUseCase(inMemoryBrightessStore, brightnessctlAdapter)
	getBrightnessPercentageUseCase := usecase.NewGetPercentageUseCase(inMemoryBrightessStore)
	persistNightLightUseCase := usecase.NewSaveUseCase(cachePersister)
	items := ui.NewListItemCollection()
	items.AddBrightness(increaseBrightnessUseCase, decreaseBrightnessUseCase, getBrightnessPercentageUseCase)

	err = startNightlightServices.Exec()
	if err != nil {
		return nil, err
	}

	if hyprsunsetAdapter.IsAvailable() {
		items.AddNightLight(increaseNightLightUseCase, decreaseNightLightUseCase, getNightLightPercentageUseCase)
	}

	listModel := ui.BuildListModel(items.List)

	return ui.NewTUI(
		listModel,
		persistNightLightUseCase,
	), nil
}
