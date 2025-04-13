package adjustable

type IAdjustable interface {
	Increase(float64)
	Decrease(float64)
	GetPercentage() float64
	GetCurrentValue() int
}
