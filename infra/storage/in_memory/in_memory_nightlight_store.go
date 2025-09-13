package in_memory_storage

import (
	"lighttui/domain/adjustable/nl"
)

type InMemoryNightlightStore struct {
	nightlight *nl.Nightlight
}

func NewInMemoryNightlightStore() *InMemoryNightlightStore {
	return &InMemoryNightlightStore{nil}
}

func (store *InMemoryNightlightStore) Fetch() *nl.Nightlight {
	return store.nightlight
}

func (store *InMemoryNightlightStore) Save(nightlight *nl.Nightlight) {
	store.nightlight = nightlight
}
