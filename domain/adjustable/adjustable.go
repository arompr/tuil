package adjustable

type IAdjustable interface {
	Increase(float64)
	Decrease(float64)
	ApplyValue(int)
	GetPercentage() float64
	GetCurrentValue() int
}
