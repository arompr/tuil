package main

import (
	"fmt"
	"lighttui/controllers"
	"lighttui/infra/hyprsunset"
	"lighttui/ui"
	"os"
)

func main() {
	// Initialize dependencies
	temperatureStore, err := hyprsunset.NewTemperatureStore()
	if err != nil {
		fmt.Println("Failed to initialize temperature store:", err)
		os.Exit(1)
	}

	brightnessCtl := controllers.NewBrightnessCtlController()
	nightLightCtl := controllers.NewNighLightController(temperatureStore)

	// Start TUI
	tui := ui.NewTUI(temperatureStore, brightnessCtl, nightLightCtl)
	if _, err := tui.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
