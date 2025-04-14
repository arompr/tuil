package cached_storage

import "lighttui/domain/adjustable"

type CachedNightLightStore struct {
	inMemoryStore adjustable.IAdjustableStore
	store         adjustable.IAdjustableStore
}

func NewCachedNightLightStore(
	inMemoryStore adjustable.IAdjustableStore,
	store adjustable.IAdjustableStore,
) *CachedNightLightStore {
	return &CachedNightLightStore{inMemoryStore, store}
}

func (c *CachedNightLightStore) Fetch() (adjustable.IAdjustable, error) {
	cached, _ := c.inMemoryStore.Fetch()
	if cached != nil {
		return cached, nil
	}

	return c.store.Fetch()
}

func (c *CachedNightLightStore) Save(adjustable adjustable.IAdjustable) error {
	return c.inMemoryStore.Save(adjustable)
}
