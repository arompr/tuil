package startup

import (
	"errors"

	"lighttui/domain/adjustable/nl"
)

type StartNightlightServices struct {
	nightlightAdapter nl.INightlightAdapter
	nightlightStore   nl.INightlightStore
}

func NewStartNightlightServices(adapter nl.INightlightAdapter, store nl.INightlightStore) *StartNightlightServices {
	return &StartNightlightServices{adapter, store}
}

func (service *StartNightlightServices) Exec() error {
	// Try to get current currentNightlight from adapter
	currentNightlight, err := service.nightlightAdapter.GetCurrentNightlight()
	if err != nil {
		var adapterUnavailable *nl.ErrNightlightAdapterUnavailable
		if errors.As(err, &adapterUnavailable) {
			return err
		}

		// fallback to store
		currentNightlight, err = service.nightlightStore.Fetch()
		if err != nil {
			return err
		}
	} else {
		// persist adapter nightlight to store
		if err := service.nightlightStore.Save(currentNightlight); err != nil {
			return err
		}
	}

	var value int
	if currentNightlight.IsEnabled() {
		value = currentNightlight.GetCurrentValue()
	} else {
		value = nl.MinTemperature
	}

	// start the adapter if needed
	if err := service.nightlightAdapter.Start(value); err != nil {
		return err
	}

	// ensure system has correct value applied
	return service.nightlightAdapter.ApplyNightlight(currentNightlight)
}
