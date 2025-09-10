package startup

import (
	"fmt"

	"lighttui/domain/adjustable/nightlight"
)

type StartNightlightServices struct {
	nightlightAdapter nightlight.INightlightAdapter
	nightlightStore   nightlight.INightlightStore
}

func NewStartNightlightServices(adapter nightlight.INightlightAdapter, store nightlight.INightlightStore) *StartNightlightServices {
	return &StartNightlightServices{adapter, store}
}

func (service *StartNightlightServices) Exec() error {
	if !service.nightlightAdapter.IsAvailable() {
		return fmt.Errorf("hyprsunset adapter is not available (is Hyprland running on this TTY?)")
	}

	// Try to get current nightlight from adapter
	nightlight, err := service.nightlightAdapter.GetCurrentNightlight()
	if err != nil || nightlight == nil {
		// fallback to store
		nightlight, err = service.nightlightStore.Fetch()
		if err != nil {
			return err
		}
	} else {
		// persist adapter nightlight to store
		if err := service.nightlightStore.Save(nightlight); err != nil {
			return err
		}
	}

	// start the adapter if needed
	if err := service.nightlightAdapter.Start(nightlight.GetCurrentValue()); err != nil {
		return err
	}

	// ensure system has correct value applied
	return service.nightlightAdapter.ApplyValue(nightlight)
}
