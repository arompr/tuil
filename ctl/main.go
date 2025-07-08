package main

import (
	"flag"
	"fmt"
	"lighttui/infra/hyprsunset"
	cached_storage "lighttui/infra/storage/cache"
	file_storage "lighttui/infra/storage/file"
	"os"
	"path/filepath"
)

func main() {
	showTemp := flag.Bool("temperature", false, "Show current night light temperature")
	flag.Parse()

	if !*showTemp {
		flag.Usage()
		return
	}

	if err := runShowTemperature(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func runShowTemperature() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	statePath := filepath.Join(homeDir, ".local/state/tuil/nightlight_temp")
	fileNightLightStore, err := file_storage.NewHyprsunsetFileStore(statePath)
	if err != nil {
		return fmt.Errorf("failed to create nightlight file store: %w", err)
	}

	cachedStore := cached_storage.NewCachedNightLightStore(
		cached_storage.NewAdjustableCache(),
		fileNightLightStore,
	)

	adapter := hyprsunset.NewNighLightAdapter(cachedStore)

	if !adapter.IsAvailable() {
		return fmt.Errorf("hyprsunset adapter is not available (is Hyprland running on this TTY?)")
	}

	if err := adapter.Start(); err != nil {
		return fmt.Errorf("failed to start hyprsunset: %w", err)
	}

	temperature, err := cachedStore.Fetch()
	if err != nil {
		return fmt.Errorf("failed to get nightlight temperature: %w", err)
	}

	fmt.Printf("Nightlight temperature: %d%%\n", int64(temperature.GetPercentage()*100))
	return nil
}
