package nightlight

type INightLightStore interface {
	Save(*NightLight) error
	Fetch() *NightLight
}

type ITemperatureStoreDeprecated interface {
	Save(int) error
	GetTemperature() int
}
