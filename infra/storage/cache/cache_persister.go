package cached_storage

type CachePersister struct {
	cachedStore *CachedNightLightStore
}

func NewCachePersister(cachedStore *CachedNightLightStore) *CachePersister {
	return &CachePersister{cachedStore}
}

func (c *CachePersister) Persist() error {
	return c.cachedStore.Persist()
}
