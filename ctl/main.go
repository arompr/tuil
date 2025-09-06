package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"lighttui/application/usecase"
	"lighttui/infra/hyprsunset"
	cached_storage "lighttui/infra/storage/cache"
	file_storage "lighttui/infra/storage/file"
)

func main() {
	showNight := flag.Bool("night", false, "Show current night light temperature")
	applyLight := flag.Bool("light", false, "Apply light temperature (6000K)")
	flag.Parse()

	if !*showNight && !*applyLight {
		flag.Usage()
		return
	}

	ctl, err := initCtl()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %v\n", err)
		os.Exit(1)
	}

	// Select and run the correct function
	if *showNight {
		if err := ctl.RunApplyNightTemperature(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	if *applyLight {
		if err := ctl.RunApplyLightTemperature(6000); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}

// Controller holds initialized dependencies and use cases
type Controller struct {
	CachedNightLightStore          *cached_storage.CachedNightLightStore
	HyprsunsetAdapter              *hyprsunset.HyprsunsetAdapter
	PersistNightLightUseCase       *usecase.SaveUseCase
	ApplyTemperatureUseCase        *usecase.ApplyTemperatureUseCase
	GetNightLightPercentageUseCase *usecase.GetPercentageUseCase
}

// initCtl initializes everything and returns a controller
func initCtl() (*Controller, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine home directory: %w", err)
	}

	fileNightLightStore, err := file_storage.NewHyprsunsetFileStore(filepath.Join(homeDir, ".local/state/tuil/nightlight_temp"))
	if err != nil {
		return nil, err
	}

	inMemoryNightLightStore := cached_storage.NewAdjustableCache()
	cachedNightLightStore := cached_storage.NewCachedNightLightStore(inMemoryNightLightStore, fileNightLightStore)
	hyprsunsetAdapter := hyprsunset.NewNighLightAdapter(cachedNightLightStore)
	cachePersister := cached_storage.NewCachePersister(cachedNightLightStore)

	return &Controller{
		CachedNightLightStore:          cachedNightLightStore,
		HyprsunsetAdapter:              hyprsunsetAdapter,
		PersistNightLightUseCase:       usecase.NewSaveUseCase(cachePersister),
		ApplyTemperatureUseCase:        usecase.NewApplyTemperatureUseCase(cachedNightLightStore, hyprsunsetAdapter),
		GetNightLightPercentageUseCase: usecase.NewGetPercentageUseCase(cachedNightLightStore),
	}, nil
}

// RunApplyNightTemperature prints the current nightlight percentage
func (c *Controller) RunApplyNightTemperature() error {
	if !c.HyprsunsetAdapter.IsAvailable() {
		return fmt.Errorf("hyprsunset adapter is not available (is Hyprland running on this TTY?)")
	}

	if err := c.HyprsunsetAdapter.Start(); err != nil {
		return fmt.Errorf("failed to start hyprsunset: %w", err)
	}

	value, err := c.CachedNightLightStore.Fetch()
	if err != nil {
		return fmt.Errorf("failed to get nightlight percentage: %w", err)
	}

	c.ApplyTemperatureUseCase.Exec(value.GetCurrentValue())

	fmt.Printf("Nightlight temperature: %d%%\n", int64(value.GetPercentage()*100))
	return nil
}

// RunApplyLightTemperature applies a light temperature and persists it
func (c *Controller) RunApplyLightTemperature(temp int) error {
	if err := c.ApplyTemperatureUseCase.Exec(temp); err != nil {
		return fmt.Errorf("failed to apply light temperature: %w", err)
	}

	if err := c.PersistNightLightUseCase.Exec(); err != nil {
		return fmt.Errorf("failed to persist light temperature: %w", err)
	}

	fmt.Printf("Light temperature applied and persisted (%dK).\n", temp)
	return nil
}
