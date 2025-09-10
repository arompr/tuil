package usecase

import "lighttui/domain/adjustable"

type IAdjustableUseCase interface {
	Exec(percentage float64) error
}

type IGetAdjustablePercentageUseCase interface {
	Exec() (float64, error)
}

func decrease(adjustable adjustable.IAdjustable, percentage float64) {
	adjustable.Decrease(percentage)
}

func increase(adjustable adjustable.IAdjustable, percentage float64) {
	adjustable.Increase(percentage)
}
