package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"lighttui/application/startup"
	"lighttui/application/usecase"
	"lighttui/infra/hyprsunset"
	cached_storage "lighttui/infra/storage/cache"
	file_storage "lighttui/infra/storage/file"
	in_memory_storage "lighttui/infra/storage/in_memory"
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

type Controller struct {
	nightlightStore                *cached_storage.CachedNightlightStore
	nightlightAdapter              *hyprsunset.HyprsunsetAdapter
	saveNightlightUseCase          *usecase.SaveUseCase
	applyTemperatureUseCase        *usecase.ApplyTemperatureUseCase
	getNightlightPercentageUseCase *usecase.GetNightlightPercentageUseCase
	turnOffNightlightUseCase       *usecase.TurnOffNightlightUseCase
	startNightlightServices        *startup.StartNightlightServices
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

	inMemoryNightLightStore := in_memory_storage.NewInMemoryNightlightStore()
	nightlightStore := cached_storage.NewCachedNightlightStore(inMemoryNightLightStore, fileNightLightStore)
	hyprsunsetAdapter := hyprsunset.NewHyprsunsetAdapter()
	cachePersister := cached_storage.NewCachePersister(nightlightStore)

	return &Controller{
		nightlightStore,
		hyprsunsetAdapter,
		usecase.NewSaveUseCase(cachePersister),
		usecase.NewApplyTemperatureUseCase(nightlightStore, hyprsunsetAdapter),
		usecase.NewGetNightlightPercentageUseCase(nightlightStore),
		usecase.NewTurnOffNightlightUseCase(nightlightStore),
		startup.NewStartNightlightServices(hyprsunsetAdapter, nightlightStore),
	}, nil
}

func (c *Controller) RunApplyNightTemperature() error {
	nightlight, err := c.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	nightlight.TurnOn()

	if err := c.nightlightStore.Save(nightlight); err != nil {
		return fmt.Errorf("failed to turn off light temperature: %w", err)
	}

	if err := c.saveNightlightUseCase.Exec(); err != nil {
		return fmt.Errorf("failed to persist: %w", err)
	}

	// start the adapter if needed
	if err := c.nightlightAdapter.Start(nightlight.GetCurrentValue()); err != nil {
		return err
	}

	// ensure system has correct value applied
	if err := c.nightlightAdapter.ApplyNightlight(nightlight); err != nil {
		return err
	}

	percentage, err := c.getNightlightPercentageUseCase.Exec()
	if err != nil {
		return err
	}

	fmt.Printf("Nightlight temperature: %d%%\n", int64(percentage*100))
	return nil
}

// RunApplyLightTemperature applies a light temperature and persists it
func (c *Controller) RunApplyLightTemperature(temp int) error {
	nightlight, err := c.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	nightlight.TurnOff()

	if err := c.nightlightAdapter.ApplValue(nightlight.GetMin()); err != nil {
		return fmt.Errorf("failed to apply light temperature: %w", err)
	}

	if err := c.nightlightStore.Save(nightlight); err != nil {
		return fmt.Errorf("failed to turn off light temperature: %w", err)
	}

	if err := c.saveNightlightUseCase.Exec(); err != nil {
		return fmt.Errorf("failed to persist: %w", err)
	}

	fmt.Printf("Nightlight turned off (%dK).\n", temp)
	return nil
}
