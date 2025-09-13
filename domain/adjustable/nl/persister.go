package nl

type IPersister interface {
	Persist() error
}
