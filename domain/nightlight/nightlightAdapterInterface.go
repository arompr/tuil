package nightlight

type INightLightAdapter interface {
	GetPercentage() float64
	IncreasePercentage(float64)
	DecreasePercentage(float64)
}
