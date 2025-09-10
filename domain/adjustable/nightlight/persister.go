package nightlight

type IPersister interface {
	Persist() error
}
