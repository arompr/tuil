package usecase

import (
	"lighttui/domain/adjustable/nl"
)

type ApplyTemperatureUseCase struct {
	nightlightStore   nl.INightlightStore
	nightlightAdapter nl.INightlightAdapter
}

func NewApplyTemperatureUseCase(
	nightlightStore nl.INightlightStore,
	nightlightAdapter nl.INightlightAdapter,
) *ApplyTemperatureUseCase {
	return &ApplyTemperatureUseCase{nightlightStore, nightlightAdapter}
}

func (usecase *ApplyTemperatureUseCase) Exec(value int) error {
	nightlight, err := usecase.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	nightlight.ApplyValue(value)

	if err := usecase.nightlightAdapter.ApplyNightlight(nightlight); err != nil {
		return err
	}

	usecase.nightlightStore.Save(nightlight)
	return nil
}
