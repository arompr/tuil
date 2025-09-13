package usecase

import (
	"lighttui/domain/adjustable"
	"lighttui/domain/adjustable/nl"
)

type AdjustNightlightUseCase struct {
	nightlightStore nl.INightlightStore
	adapter         nl.INightlightAdapter
	adjust          func(adjustable.IAdjustable, float64)
}

func NewDecreaseNightlightUseCase(
	nightlightStore nl.INightlightStore,
	adapter nl.INightlightAdapter,
) *AdjustNightlightUseCase {
	return &AdjustNightlightUseCase{nightlightStore, adapter, decrease}
}

func NewIncreaseNightlightUseCase(
	nightlightStore nl.INightlightStore,
	adapter nl.INightlightAdapter,
) *AdjustNightlightUseCase {
	return &AdjustNightlightUseCase{nightlightStore, adapter, increase}
}

func (usecase *AdjustNightlightUseCase) Exec(percentage float64) error {
	nightlight, err := usecase.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	usecase.adjust(nightlight, percentage)

	if err := usecase.adapter.ApplyNightlight(nightlight); err != nil {
		return err
	}

	usecase.nightlightStore.Save(nightlight)
	return nil
}
