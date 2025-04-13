package usecase

import (
	"lighttui/domain/adjustable"
)

type PersistUseCase struct {
	persistantStore adjustable.IPersistentAdjustableStore
	store           adjustable.IAdjustableStore
}

func NewPersistUseCase(
	persistantStore adjustable.IPersistentAdjustableStore,
	store adjustable.IAdjustableStore,
) *PersistUseCase {
	return &PersistUseCase{persistantStore, store}
}

func (p *PersistUseCase) Exec() error {
	return p.persistantStore.Save(p.store.Fetch())
}
