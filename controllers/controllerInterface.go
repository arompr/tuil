package controllers

// BrightnessController defines the required methods for brightness control.
type IController interface {
	GetPercentage() float64
	GetCurrent() int
	IncreasePercentage(float64)
	DecreasePercentage(float64)
}
