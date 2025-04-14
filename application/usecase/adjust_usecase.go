package usecase

import (
	"lighttui/domain/adjustable"
)

type AdjustUseCase struct {
	store   adjustable.IAdjustableStore
	adapter adjustable.IAdjustableAdapter
	adjust  func(adjustable.IAdjustable, float64)
}

func NewDecreaseUseCase(
	store adjustable.IAdjustableStore,
	adapter adjustable.IAdjustableAdapter,
) *AdjustUseCase {
	return &AdjustUseCase{store, adapter, decrease}
}

func decrease(i adjustable.IAdjustable, percentage float64) {
	i.Decrease(percentage)
}

func NewIncreaseUseCase(
	store adjustable.IAdjustableStore,
	adapter adjustable.IAdjustableAdapter,
) *AdjustUseCase {
	return &AdjustUseCase{store, adapter, increase}
}

func increase(i adjustable.IAdjustable, percentage float64) {
	i.Increase(percentage)
}

func (i *AdjustUseCase) Exec(percentage float64) error {
	adjustable, _ := i.store.Fetch()
	i.adjust(adjustable, percentage)

	if err := i.adapter.ApplyValue(adjustable); err != nil {
		return err
	}

	i.store.Save(adjustable)
	return nil
}
