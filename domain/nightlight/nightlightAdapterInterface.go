package nightlight

type INightLightAdapter interface {
	ApplyNightLight(nightlight *NightLight) error
}
