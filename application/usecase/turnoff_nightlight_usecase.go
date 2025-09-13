package usecase

import "lighttui/domain/adjustable/nl"

type TurnOffNightlightUseCase struct {
	nightlightStore nl.INightlightStore
}

func NewTurnOffNightlightUseCase(
	nightlightStore nl.INightlightStore,
) *TurnOffNightlightUseCase {
	return &TurnOffNightlightUseCase{nightlightStore}
}

func (usecase *TurnOffNightlightUseCase) Exec() error {
	nightlight, err := usecase.nightlightStore.Fetch()
	if err != nil {
		return err
	}

	nightlight.TurnOff()
	usecase.nightlightStore.Save(nightlight)

	return nil
}
