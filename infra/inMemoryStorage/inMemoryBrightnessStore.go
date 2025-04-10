package in_memory_storage

import (
	"lighttui/domain/brightness"
)

type InMemoryBrightnessStore struct {
	brightness *brightness.Brightness
}

func NewInMemoryBrightnessStore(brightness *brightness.Brightness) *InMemoryBrightnessStore {
	return &InMemoryBrightnessStore{brightness}
}

func (s *InMemoryBrightnessStore) FetchBrightness() *brightness.Brightness {
	return s.brightness
}

func (s *InMemoryBrightnessStore) Save(brightness *brightness.Brightness) {
	s.brightness = brightness
}
