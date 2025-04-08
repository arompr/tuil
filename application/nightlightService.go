package service

import (
	"lighttui/domain/nightlight"
)

type NightLightService struct {
	store   nightlight.INightLightStore
	gateway nightlight.INightLightAdapter
}

func NewNightLightService(
	store nightlight.INightLightStore,
	gateway nightlight.INightLightAdapter,
) *NightLightService {
	return &NightLightService{store, gateway}
}

func (s *NightLightService) Increase(percentage float64) error {
	nightlight := s.store.FetchNightLight()
	nightlight.Increase(percentage)
	return s.applyNightLight(nightlight)
}

func (s *NightLightService) Decrease(percentage float64) error {
	nightlight := s.store.FetchNightLight()
	nightlight.Decrease(percentage)
	return s.applyNightLight(nightlight)
}

func (s *NightLightService) applyNightLight(nightlight *nightlight.NightLight) error {
	if err := s.gateway.ApplyNightLight(nightlight); err != nil {
		return err
	}

	s.store.Save(nightlight)
	return nil
}
