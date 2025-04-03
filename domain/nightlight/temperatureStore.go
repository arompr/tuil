package nightlight

type ITemperatureStore interface {
	Save(NightLight) error
	GetTemperature() NightLight
}

type ITemperatureStoreDeprecated interface {
	Save(int) error
	GetTemperature() int
}
