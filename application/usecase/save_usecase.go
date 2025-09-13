package usecase

import "lighttui/domain/adjustable/nl"

type SaveUseCase struct {
	persister nl.IPersister
}

func NewSaveUseCase(nightlightPersister nl.IPersister) *SaveUseCase {
	return &SaveUseCase{nightlightPersister}
}

func (usecase *SaveUseCase) Exec() error {
	return usecase.persister.Persist()
}
