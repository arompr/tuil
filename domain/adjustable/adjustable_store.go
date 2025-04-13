package adjustable

type IAdjustableStore interface {
	Save(IAdjustable)
	Fetch() IAdjustable
}

type IPersistentAdjustableStore interface {
	Save(IAdjustable) error
	Fetch() (IAdjustable, error)
}
