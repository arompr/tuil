package service

import (
	"lighttui/domain/brightness"
)

type BrightnessService struct {
	store   brightness.IBrightnessStore
	gateway brightness.IBrightnessAdapter
}

func NewBrightnessService(
	store brightness.IBrightnessStore,
	gateway brightness.IBrightnessAdapter,
) *BrightnessService {
	return &BrightnessService{store, gateway}
}

func (s *BrightnessService) Increase(percentage float64) error {
	brightness := s.store.FetchBrightness()
	brightness.Increase(percentage)
	return s.applyNightLight(brightness)
}

func (s *BrightnessService) Decrease(percentage float64) error {
	nightlight := s.store.FetchBrightness()
	nightlight.Decrease(percentage)
	return s.applyNightLight(nightlight)
}

func (s *BrightnessService) applyNightLight(brightness *brightness.Brightness) error {
	if err := s.gateway.ApplyBrightness(brightness); err != nil {
		return err
	}

	s.store.Save(brightness)
	return nil
}
