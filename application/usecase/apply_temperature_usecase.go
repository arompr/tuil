package usecase

import (
	"lighttui/domain/adjustable/nightlight"
)

type ApplyTemperatureUseCase struct {
	nightlightStore   nightlight.INightlightStore
	nightlightAdapter nightlight.INightlightAdapter
}

func NewApplyTemperatureUseCase(
	nightlightStore nightlight.INightlightStore,
	nightlightAdapter nightlight.INightlightAdapter,
) *ApplyTemperatureUseCase {
	return &ApplyTemperatureUseCase{nightlightStore, nightlightAdapter}
}

func (usecase *ApplyTemperatureUseCase) Exec(value int) error {
	nightlight, err := usecase.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	nightlight.ApplyValue(value)

	if err := usecase.nightlightAdapter.ApplyValue(nightlight); err != nil {
		return err
	}

	usecase.nightlightStore.Save(nightlight)
	return nil
}
