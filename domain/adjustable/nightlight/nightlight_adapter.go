package nightlight

import "lighttui/domain/adjustable"

type INightLightAdapter interface {
	ApplyValue(adjustable.IAdjustable) error
}
