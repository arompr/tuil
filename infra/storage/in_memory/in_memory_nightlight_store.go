package in_memory_storage

import (
	"lighttui/domain/adjustable"
)

type InMemoryNightLightStore struct {
	nightlight adjustable.IAdjustable
}

func NewInMemoryNightLightStore() *InMemoryNightLightStore {
	return &InMemoryNightLightStore{nil}
}

func (s *InMemoryNightLightStore) Fetch() adjustable.IAdjustable {
	return s.nightlight
}

// Save stores the NightLight in memory. Always returns nil.
func (s *InMemoryNightLightStore) Save(adjustable adjustable.IAdjustable) {
	s.nightlight = adjustable
}
