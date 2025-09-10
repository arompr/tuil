package nightlight

type Nightlight struct {
	currentTemperature int
	max                int
	min                int
	enabled            bool
}

const (
	MaxTemperature = 1500
	MinTemperature = 6000
)

type NightlightOption func(*Nightlight)

func WithEnabled(enabled bool) NightlightOption {
	return func(n *Nightlight) {
		n.enabled = enabled
	}
}

func CreateNewNightLight(value int, opts ...NightlightOption) *Nightlight {
	n := &Nightlight{value, MaxTemperature, MinTemperature, true}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func (n *Nightlight) Increase(percentage float64) {
	n.ApplyValue(max(n.calculateNewTemperature(-percentage), n.GetMax()))
}

func (n *Nightlight) Decrease(percentage float64) {
	n.ApplyValue(min(n.calculateNewTemperature(percentage), n.GetMin()))
}

func (n *Nightlight) ApplyValue(value int) {
	n.currentTemperature = value
}

func (n *Nightlight) calculateNewTemperature(percentage float64) int {
	return int(float64(n.GetCurrentValue()) + n.getDelta(percentage))
}

// delta is the amount of temperature change that corresponds to a given percentage of the full night light range.
func (n *Nightlight) getDelta(percentage float64) float64 {
	return float64(n.GetMin()-n.GetMax()) * percentage
}

func (n *Nightlight) GetPercentage() float64 {
	return 1 - (float64(n.GetCurrentValue()-n.GetMax()) / float64(n.GetMin()-n.GetMax()))
}

func (n *Nightlight) GetCurrentValue() int {
	return n.currentTemperature
}

func (n *Nightlight) GetMax() int {
	return n.max
}

func (n *Nightlight) GetMin() int {
	return n.min
}

func (n *Nightlight) IsEnabled() bool {
	return n.enabled
}
