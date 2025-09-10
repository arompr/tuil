package cached_storage

import (
	"lighttui/domain/adjustable/nightlight"
	file_storage "lighttui/infra/storage/file"
	in_memory_storage "lighttui/infra/storage/in_memory"
)

type CachedNightlightStore struct {
	inMemoryNightlightStore *in_memory_storage.InMemoryNightlightStore
	store                   *file_storage.FileNightlightStore
}

func NewCachedNightlightStore(
	inMemoryNightlightStore *in_memory_storage.InMemoryNightlightStore,
	store *file_storage.FileNightlightStore,
) *CachedNightlightStore {
	return &CachedNightlightStore{inMemoryNightlightStore, store}
}

func (store *CachedNightlightStore) Fetch() (*nightlight.Nightlight, error) {
	inMemoryNightlight := store.inMemoryNightlightStore.Fetch()
	if inMemoryNightlight != nil {
		return inMemoryNightlight, nil
	}

	persistedNightlight, err := store.store.Fetch()
	if err != nil {
		return nil, err
	}

	store.inMemoryNightlightStore.Save(persistedNightlight)
	return persistedNightlight, nil
}

func (store *CachedNightlightStore) Save(nightlight *nightlight.Nightlight) error {
	store.inMemoryNightlightStore.Save(nightlight)
	return nil
}

func (store *CachedNightlightStore) Persist() error {
	nightlight, err := store.Fetch()
	if err != nil {
		return err
	}

	return store.store.Save(nightlight)
}
