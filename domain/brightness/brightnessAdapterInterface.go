package brightness

type IBrightnessAdapter interface {
	GetPercentage() float64
	IncreasePercentage(float64)
	DecreasePercentage(float64)
}
