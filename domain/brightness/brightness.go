package brightness

type Brightness struct {
	currentBrightness int
	max               int
	min               int
}

const MinBrightness = 1

func CreateNewBrightness(value, max int) *Brightness {
	return &Brightness{value, max, MinBrightness}
}

func (b *Brightness) Increase(percentage float64) {
	b.applyBrightness(int(min(b.calculateNewBrightness(percentage), float64(b.GetMax()))))
}

func (b *Brightness) Decrease(percentage float64) {
	b.applyBrightness(int(max(b.calculateNewBrightness(-percentage), float64(b.GetMin()))))
}

func (b *Brightness) calculateNewBrightness(percentage float64) float64 {
	return float64(b.GetCurrentValue()) + (float64(b.GetMax()) * percentage)
}

func (b *Brightness) applyBrightness(value int) {
	b.currentBrightness = value
}

func (b *Brightness) GetPercentage() float64 {
	return float64(b.currentBrightness) / float64(b.GetMax())
}

func (b *Brightness) GetCurrentValue() int {
	return b.currentBrightness
}

func (b *Brightness) GetMax() int {
	return b.max
}

func (b *Brightness) GetMin() int {
	return b.min
}
