package cached_storage

import (
	"lighttui/domain/adjustable"
)

type AdjustableCache struct {
	adjustable adjustable.IAdjustable
}

func NewAdjustableCache() *AdjustableCache {
	return &AdjustableCache{nil}
}

func (a *AdjustableCache) Fetch() adjustable.IAdjustable {
	return a.adjustable
}

func (a *AdjustableCache) Save(adjustable adjustable.IAdjustable) {
	a.adjustable = adjustable
}
