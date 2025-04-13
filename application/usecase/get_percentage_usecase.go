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

func (g *GetPercentageUseCase) Exec() float64 {
	return g.store.Fetch().GetPercentage()
}
