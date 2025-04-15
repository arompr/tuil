package usecase

type SaveUseCase struct {
	persister IPersister
}

func NewSaveUseCase(cacheStorecacheUnitOfWork IPersister) *SaveUseCase {
	return &SaveUseCase{cacheStorecacheUnitOfWork}
}

func (p *SaveUseCase) Exec() error {
	return p.persister.Persist()
}
