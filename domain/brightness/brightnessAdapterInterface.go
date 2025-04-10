package brightness

type IBrightnessAdapter interface {
	GetCurrentBrightnessValue() int
	GetMaxBrightnessValue() int
	ApplyBrightness(brightness *Brightness) error
}
