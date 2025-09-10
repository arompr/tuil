package usecase

import (
	"lighttui/domain/adjustable"
	"lighttui/domain/adjustable/nightlight"
)

type AdjustNightlightUseCase struct {
	nightlightStore nightlight.INightlightStore
	adapter         nightlight.INightlightAdapter
	adjust          func(adjustable.IAdjustable, float64)
}

func NewDecreaseNightlightUseCase(
	nightlightStore nightlight.INightlightStore,
	adapter nightlight.INightlightAdapter,
) *AdjustNightlightUseCase {
	return &AdjustNightlightUseCase{nightlightStore, adapter, decrease}
}

func NewIncreaseNightlightUseCase(
	nightlightStore nightlight.INightlightStore,
	adapter nightlight.INightlightAdapter,
) *AdjustNightlightUseCase {
	return &AdjustNightlightUseCase{nightlightStore, adapter, increase}
}

func (usecase *AdjustNightlightUseCase) Exec(percentage float64) error {
	nightlight, err := usecase.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	usecase.adjust(nightlight, percentage)

	if err := usecase.adapter.ApplyValue(nightlight); err != nil {
		return err
	}

	usecase.nightlightStore.Save(nightlight)
	return nil
}
