package usecase

import (
	"lighttui/domain/adjustable"
)

type GetPercentageUseCase struct {
	store adjustable.IAdjustableStore
}

func NewGetPercentageUseCase(store adjustable.IAdjustableStore) *GetPercentageUseCase {
	return &GetPercentageUseCase{store}
}

func (g *GetPercentageUseCase) Exec() (float64, error) {
	adjustable, err := g.store.Fetch()
	if err != nil {
		return 0, err
	}

	return adjustable.GetPercentage(), nil
}
