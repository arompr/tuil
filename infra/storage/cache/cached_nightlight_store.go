package cached_storage

import (
	"lighttui/domain/adjustable"
)

type CachedNightLightStore struct {
	cache *AdjustableCache
	store adjustable.IAdjustableStore
}

func NewCachedNightLightStore(
	cache *AdjustableCache,
	store adjustable.IAdjustableStore,
) *CachedNightLightStore {
	return &CachedNightLightStore{cache, store}
}

func (c *CachedNightLightStore) Fetch() (adjustable.IAdjustable, error) {
	cached := c.cache.Fetch()
	if cached != nil {
		return cached, nil
	}

	adjustable, err := c.store.Fetch()
	if err != nil {
		return nil, err
	}

	c.cache.Save(adjustable)
	return adjustable, nil
}

func (c *CachedNightLightStore) Save(adjustable adjustable.IAdjustable) error {
	c.cache.Save(adjustable)
	return nil
}
