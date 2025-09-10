package usecase

import "lighttui/domain/adjustable/nightlight"

type SaveUseCase struct {
	persister nightlight.IPersister
}

func NewSaveUseCase(cacheStorecacheUnitOfWork nightlight.IPersister) *SaveUseCase {
	return &SaveUseCase{cacheStorecacheUnitOfWork}
}

func (p *SaveUseCase) Exec() error {
	return p.persister.Persist()
}
