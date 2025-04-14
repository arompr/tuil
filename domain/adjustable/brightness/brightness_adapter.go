package brightness

import "lighttui/domain/adjustable"

type IBrightnessAdapter interface {
	GetCurrentBrightnessValue() (int, error)
	GetMaxBrightnessValue() (int, error)
	ApplyValue(adjustable.IAdjustable) error
}
