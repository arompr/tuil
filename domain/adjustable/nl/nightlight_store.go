package nl

type INightlightStore interface {
	Save(*Nightlight) error
	Fetch() (*Nightlight, error)
}
