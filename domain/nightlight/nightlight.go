package nightlight

type NightLight struct {
	currentTemperature int
	max                int
	min                int
}

const MaxTemperature = 1500
const MinTemperature = 6500

func CreateNewNightLight(value int) *NightLight {
	return &NightLight{value, MaxTemperature, MinTemperature}
}

func (n *NightLight) Increase(percentage float64) {
	n.applyTemperature(max(n.calculateNewTemperature(-percentage), n.GetMax()))
}

func (n *NightLight) Decrease(percentage float64) {
	n.applyTemperature(min(n.calculateNewTemperature(percentage), n.GetMin()))
}

func (n *NightLight) applyTemperature(value int) {
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
