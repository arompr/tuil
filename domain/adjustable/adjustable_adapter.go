package adjustable

type IAdjustableAdapter interface {
	ApplyValue(IAdjustable) error
}
