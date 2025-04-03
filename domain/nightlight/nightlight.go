package nightlight

const maxTemperature = 1500
const minTemperature = 6500

type NightLight struct {
	value int
}

func CreateNewNightLight(value int) *NightLight {
	return &NightLight{value: value}
}

func (n *NightLight) Increase(percentage int) {
	n.setValue(n.calculateNewValue(1-percentage, maxTemperature))
}

func (n *NightLight) Decrease(percentage int) {
	n.setValue(n.calculateNewValue(1+percentage, minTemperature))
}

func (n *NightLight) setValue(newValue int) {
	n.value = newValue
}

func (n *NightLight) calculateNewValue(ratio, other int) int {
	return max(int(float64(n.value)*float64(ratio)), other)
}
