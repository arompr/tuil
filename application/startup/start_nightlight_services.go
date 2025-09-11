package startup

import (
	"errors"

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
	// Try to get current currentNightlight from adapter
	currentNightlight, err := service.nightlightAdapter.GetCurrentNightlight()
	if err != nil {
		var adapterUnavailable *nightlight.ErrNightlightAdapterUnavailable
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

	// start the adapter if needed
	if err := service.nightlightAdapter.Start(currentNightlight.GetCurrentValue()); err != nil {
		return err
	}

	// ensure system has correct value applied
	return service.nightlightAdapter.ApplyValue(currentNightlight)
}
