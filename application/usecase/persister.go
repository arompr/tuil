package usecase

type IPersister interface {
	Persist() error
}
