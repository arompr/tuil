package service

import (
	"lighttui/domain/adjustable"
	"lighttui/domain/nightlight"
)

type NightLightService struct {
	store   adjustable.IAdjustableStore
	gateway nightlight.INightLightAdapter
}

func NewNightLightService(
	store adjustable.IAdjustableStore,
	gateway nightlight.INightLightAdapter,
) *NightLightService {
	return &NightLightService{store, gateway}
}

func (s *NightLightService) Increase(percentage float64) error {
	nightlight := s.store.Fetch()
	nightlight.Increase(percentage)
	return s.applyNightLight(nightlight)
}

func (s *NightLightService) Decrease(percentage float64) error {
	nightlight := s.store.Fetch()
	nightlight.Decrease(percentage)
	return s.applyNightLight(nightlight)
}

func (s *NightLightService) applyNightLight(nightlight adjustable.IAdjustable) error {
	if err := s.gateway.ApplyValue(nightlight); err != nil {
		return err
	}

	s.store.Save(nightlight)
	return nil
}
