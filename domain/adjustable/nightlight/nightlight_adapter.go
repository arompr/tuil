package nightlight

import "lighttui/domain/adjustable"

type INightLightAdapter interface {
	IsAvailable() bool
	Start() error
	ApplyValue(adjustable.IAdjustable) error
}
