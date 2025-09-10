package usecase

import (
	"lighttui/domain/adjustable/nightlight"
)

type GetNightlightPercentageUseCase struct {
	store nightlight.INightlightStore
}

func NewGetNightlightPercentageUseCase(store nightlight.INightlightStore) *GetNightlightPercentageUseCase {
	return &GetNightlightPercentageUseCase{store}
}

func (usecase *GetNightlightPercentageUseCase) Exec() (float64, error) {
	nightlight, err := usecase.store.Fetch()
	if err != nil {
		return 0, err
	}

	return nightlight.GetPercentage(), nil
}
