package nightlight

type NightLight struct {
	currentTemperature int
	max                int
	min                int
}

func CreateNewNightLight(value, max, min int) *NightLight {
	return &NightLight{value, max, min}
}

func (n *NightLight) Increase(percentage float64) {
	n.applyTemperature(n.calculateNewTemperature(1-percentage, n.GetMax()))
}

func (n *NightLight) Decrease(percentage float64) {
	n.applyTemperature(n.calculateNewTemperature(1+percentage, n.GetMin()))
}

func (n *NightLight) applyTemperature(value int) {
	n.currentTemperature = value
}

func (n *NightLight) calculateNewTemperature(ratio float64, other int) int {
	return max(int(float64(n.GetCurrentTemperature())*ratio), other)
}

func (n *NightLight) GetPercentage() float64 {
	return 1 - (float64(n.GetCurrentTemperature()-n.GetMax()) / float64(n.GetMin()-n.GetMax()))
}

func (n *NightLight) GetCurrentTemperature() int {
	return n.currentTemperature
}

func (n *NightLight) GetMax() int {
	return n.max
}

func (n *NightLight) GetMin() int {
	return n.min
}
