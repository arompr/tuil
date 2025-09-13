package nl

type INightlightAdapter interface {
	IsAvailable() bool
	Start(value int) error
	ApplyNightlight(*Nightlight) error
	GetCurrentNightlight() (*Nightlight, error)
}
