package brightness

type IBrightnessStore interface {
	Save(*Brightness)
	FetchBrightness() *Brightness
}
