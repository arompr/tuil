package adjustable

type IAdjustableStore interface {
	Save(IAdjustable) error
	Fetch() IAdjustable
}
