package nightlight

type INightlightStore interface {
	Save(*Nightlight) error
	Fetch() (*Nightlight, error)
}
