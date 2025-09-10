package in_memory_storage

import (
	"lighttui/domain/adjustable/nightlight"
)

type InMemoryNightlightStore struct {
	nightlight *nightlight.Nightlight
}

func NewInMemoryNightlightStore() *InMemoryNightlightStore {
	return &InMemoryNightlightStore{nil}
}

func (store *InMemoryNightlightStore) Fetch() *nightlight.Nightlight {
	return store.nightlight
}

func (store *InMemoryNightlightStore) Save(nightlight *nightlight.Nightlight) {
	store.nightlight = nightlight
}
