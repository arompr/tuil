package brightness

type Brightness struct {
	currentBrightness int
	max               int
	min               int
}

func CreateNewNightLight(value, max, min int) *Brightness {
	return &Brightness{value, max, min}
}

func (b *Brightness) Increase(percentage float64) {
	b.applyBrightness(int(min(float64(b.GetCurrentBrightness())+(float64(b.GetMax())*percentage), float64(b.GetMax()))))
}

func (b *Brightness) Decrease(percentage float64) {
	b.applyBrightness(int(max(float64(b.GetCurrentBrightness())-(float64(b.GetMax())*percentage), float64(b.GetMin()))))
}

func (b *Brightness) applyBrightness(value int) {
	b.currentBrightness = value
}

func (b *Brightness) GetPercentage() float64 {
	return float64(b.currentBrightness) / float64(b.GetMax())
}

func (b *Brightness) GetCurrentBrightness() int {
	return b.currentBrightness
}

func (b *Brightness) GetMax() int {
	return b.max
}

func (b *Brightness) GetMin() int {
	return b.min
}
