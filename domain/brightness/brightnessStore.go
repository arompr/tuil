package brightness

type IBrightnessStore interface {
	Save(*Brightness) error
	FetchBrightness() *Brightness
}
