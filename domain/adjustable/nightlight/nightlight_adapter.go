package nightlight

type INightlightAdapter interface {
	IsAvailable() bool
	Start(value int) error
	ApplyValue(*Nightlight) error
	GetCurrentNightlight() (*Nightlight, error)
}
