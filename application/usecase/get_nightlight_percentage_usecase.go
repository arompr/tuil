package usecase

import (
	"lighttui/domain/adjustable/nl"
)

type GetNightlightPercentageUseCase struct {
	store nl.INightlightStore
}

func NewGetNightlightPercentageUseCase(store nl.INightlightStore) *GetNightlightPercentageUseCase {
	return &GetNightlightPercentageUseCase{store}
}

func (usecase *GetNightlightPercentageUseCase) Exec() (float64, error) {
	nightlight, err := usecase.store.Fetch()
	if err != nil {
		return 0, err
	}

	return nightlight.GetPercentage(), nil
}
