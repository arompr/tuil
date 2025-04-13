package main

import (
	"fmt"
	service "lighttui/application"
	"lighttui/application/usecase"
	"lighttui/controllers"
	"lighttui/domain/brightness"
	"lighttui/infra/brightnessctl"
	"lighttui/infra/hyprsunset"
	file_storage "lighttui/infra/storage/file"
	in_memory_storage "lighttui/infra/storage/in_memory"
	"lighttui/ui"
	"lighttui/ui/deprecated"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize dependencies DEP
	temperatureStore, err := hyprsunset.NewTemperatureStore()
	if err != nil {
		fmt.Println("Failed to initialize temperature store:", err)
		os.Exit(1)
	}

	brightnessCtl := controllers.NewBrightnessCtlController()
	nightLightCtl := controllers.NewNighLightController(temperatureStore)

	// Start TUI
	tui := deprecated.NewTUIDeprecated(temperatureStore, brightnessCtl, nightLightCtl)
	if _, err := tui.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

func InitDep() (*tea.Program, error) {
	// Initialise dependencies
	fileNightLightStore, err := file_storage.NewHyprsunsetFileStore("sd")
	if err != nil {
		return nil, err
	}

	nightlight, err := fileNightLightStore.FetchNightLight()
	if err != nil {
		return nil, err
	}
	inMemoryNightLightStore := in_memory_storage.NewInMemoryNightLightStore()
	inMemoryNightLightStore.Save(nightlight)

	hyprsunsetAdapter := hyprsunset.NewNighLightAdapter(inMemoryNightLightStore)
	service.NewNightLightService(inMemoryNightLightStore, hyprsunsetAdapter)

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

	increaseNightLightUseCase := usecase.NewIncreaseUseCase(inMemoryNightLightStore, hyprsunsetAdapter)
	decreaseNightLightUseCase := usecase.NewDecreaseUseCase(inMemoryNightLightStore, hyprsunsetAdapter)
	getNightLightPercentageUseCase := usecase.NewGetPercentageUseCase(inMemoryNightLightStore)
	increaseBrightnessUseCase := usecase.NewIncreaseUseCase(inMemoryBrightessStore, brightnessctlAdapter)
	decreaseBrightnessUseCase := usecase.NewDecreaseUseCase(inMemoryBrightessStore, brightnessctlAdapter)
	getBrightnessPercentageUseCase := usecase.NewGetPercentageUseCase(inMemoryBrightessStore)

	return ui.NewTUI(
		increaseNightLightUseCase,
		decreaseNightLightUseCase,
		getNightLightPercentageUseCase,
		increaseBrightnessUseCase,
		decreaseBrightnessUseCase,
		getBrightnessPercentageUseCase,
	), nil
}
