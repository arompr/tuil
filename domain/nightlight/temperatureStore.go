package nightlight

type INightLightStore interface {
	Save(*NightLight) error
	FetchNightLight() *NightLight
}

type ITemperatureStoreDeprecated interface {
	Save(int) error
	GetTemperature() int
}
