package cached_storage

type CachePersister struct {
	cachedStore *CachedNightlightStore
}

func NewCachePersister(cachedStore *CachedNightlightStore) *CachePersister {
	return &CachePersister{cachedStore}
}

func (c *CachePersister) Persist() error {
	return c.cachedStore.Persist()
}
