package usecase

import (
	"lighttui/domain/adjustable"
)

type ApplyTemperatureUseCase struct {
	store   adjustable.IAdjustableStore
	adapter adjustable.IAdjustableAdapter
}

func NewApplyTemperatureUseCase(
	store adjustable.IAdjustableStore,
	adapter adjustable.IAdjustableAdapter,
) *ApplyTemperatureUseCase {
	return &ApplyTemperatureUseCase{store, adapter}
}

func (i *ApplyTemperatureUseCase) Exec(value int) error {
	adjustable, err := i.store.Fetch()
	if err != nil {
		return err
	}

	adjustable.ApplyValue(value)

	if err := i.adapter.ApplyValue(adjustable); err != nil {
		return err
	}

	i.store.Save(adjustable)
	return nil
}
