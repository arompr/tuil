package nl

type Nightlight struct {
	currentTemperature int
	enabled            bool
}

const (
	MaxTemperature = 1500
	MinTemperature = 6000
)

type NightlightOption func(*Nightlight)

func CreateNewNightlight(value int) *Nightlight {
	var isEnabled bool
	if value == MinTemperature {
		isEnabled = false
	} else {
		isEnabled = true
	}

	return &Nightlight{value, isEnabled}
}

func (n *Nightlight) TurnOn() {
	n.enabled = true
}

func (n *Nightlight) TurnOff() {
	n.enabled = false
}

func (n *Nightlight) Increase(percentage float64) {
	n.ApplyValue(max(n.calculateNewTemperature(-percentage), n.GetMax()))
}

func (n *Nightlight) Decrease(percentage float64) {
	n.ApplyValue(min(n.calculateNewTemperature(percentage), n.GetMin()))
}

func (n *Nightlight) ApplyValue(value int) {
	if value == MinTemperature {
		n.enabled = false
	} else {
		n.enabled = true
	}

	n.currentTemperature = value
}

func (n *Nightlight) calculateNewTemperature(percentage float64) int {
	return int(float64(n.GetCurrentValue()) + n.getDelta(percentage))
}

func (n *Nightlight) GetCurrentValue() int {
	return n.currentTemperature
}

// delta is the amount of temperature change that corresponds to a given percentage of the full night light range.
func (n *Nightlight) getDelta(percentage float64) float64 {
	return float64(n.GetMin()-n.GetMax()) * percentage
}

func (n *Nightlight) GetPercentage() float64 {
	return 1 - (float64(n.GetCurrentValue()-n.GetMax()) / float64(n.GetMin()-n.GetMax()))
}

func (n *Nightlight) GetMax() int {
	return MaxTemperature
}

func (n *Nightlight) GetMin() int {
	return MinTemperature
}

func (n *Nightlight) IsEnabled() bool {
	return n.enabled
}
