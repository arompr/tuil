package usecase

import (
	"lighttui/domain/adjustable"
)

type PersistUseCase struct {
	cacheStore      adjustable.IAdjustableStore
	persistentStore adjustable.IAdjustableStore
}

func NewPersistUseCase(
	cacheStore adjustable.IAdjustableStore,
	persistentStore adjustable.IAdjustableStore,
) *PersistUseCase {
	return &PersistUseCase{cacheStore, persistentStore}
}

func (p *PersistUseCase) Exec() error {
	adjustable, err := p.cacheStore.Fetch()
	if err != nil {
		return err
	}

	return p.persistentStore.Save(adjustable)
}
