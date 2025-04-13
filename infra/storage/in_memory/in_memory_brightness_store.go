package in_memory_storage

import (
	"lighttui/domain/adjustable"
)

type InMemoryBrightnessStore struct {
	brightness adjustable.IAdjustable
}

func NewInMemoryBrightnessStore() *InMemoryBrightnessStore {
	return &InMemoryBrightnessStore{nil}
}

func (s *InMemoryBrightnessStore) Fetch() adjustable.IAdjustable {
	return s.brightness
}

func (s *InMemoryBrightnessStore) Save(adjustable adjustable.IAdjustable) {
	s.brightness = adjustable
}
