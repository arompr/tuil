package store

import (
	"lighttui/domain/nightlight"
)

type InMemoryNightLightStore struct {
	nightlight *nightlight.NightLight
}

func NewInMemoryNightLightStore(nightlight *nightlight.NightLight) *InMemoryNightLightStore {
	return &InMemoryNightLightStore{nightlight}
}

func (s *InMemoryNightLightStore) Fetch() *nightlight.NightLight {
	return s.nightlight
}

func (s *InMemoryNightLightStore) Save(nightlight *nightlight.NightLight) {
	s.nightlight = nightlight
}
