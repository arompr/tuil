package nightlight

type INightLight interface {
	IsEnabled() bool
	GetCurrentTemperature()
}

type NightLight struct {
	currentTemperature int
	max                int
	min                int
	enabled            bool
}

const (
	MaxTemperature = 1500
	MinTemperature = 6000
)

type NightLightOption func(*NightLight)

func WithEnabled(enabled bool) NightLightOption {
	return func(n *NightLight) {
		n.enabled = enabled
	}
}

func CreateNewNightLight(value int, opts ...NightLightOption) *NightLight {
	n := &NightLight{value, MaxTemperature, MinTemperature, true}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func (n *NightLight) Increase(percentage float64) {
	n.ApplyValue(max(n.calculateNewTemperature(-percentage), n.GetMax()))
}

func (n *NightLight) Decrease(percentage float64) {
	n.ApplyValue(min(n.calculateNewTemperature(percentage), n.GetMin()))
}

func (n *NightLight) ApplyValue(value int) {
	n.currentTemperature = value
}

func (n *NightLight) calculateNewTemperature(percentage float64) int {
	return int(float64(n.GetCurrentValue()) + n.getDelta(percentage))
}

// delta is the amount of temperature change that corresponds to a given percentage of the full night light range.
func (n *NightLight) getDelta(percentage float64) float64 {
	return float64(n.GetMin()-n.GetMax()) * percentage
}

func (n *NightLight) GetPercentage() float64 {
	return 1 - (float64(n.GetCurrentValue()-n.GetMax()) / float64(n.GetMin()-n.GetMax()))
}

func (n *NightLight) GetCurrentValue() int {
	return n.currentTemperature
}

func (n *NightLight) GetMax() int {
	return n.max
}

func (n *NightLight) GetMin() int {
	return n.min
}
