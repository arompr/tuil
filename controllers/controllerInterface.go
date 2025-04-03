package controllers

type IController interface {
	GetPercentage() float64
	GetCurrent() int
	IncreasePercentage(float64)
	DecreasePercentage(float64)
}
