package cache

import "time"

type Cache interface {
	Get(key string) interface{}

	Put(key string, value interface{}) interface{}

	Delete(key string)

	Size() int

	CompareAndSwap(key string, old, new interface{}) (interface{}, bool)
}

// Options control the behavior of the cache
type Options struct {
	// TTL controls the time-to-live for a given cache entry.  Cache entries that
	// are older than the TTL will not be returned
	TTL time.Duration

	// InitialCapacity controls the initial capacity of the cache
	InitialCapacity int

	// OnEvict is an optional function called when an element is evicted.
	OnEvict EvictCallback

	// TimeNow is used to override the behavior of default time.Now(), e.g. in tests.
	TimeNow func() time.Time
}

// EvictCallback is a type for notifying applications when an item is
// scheduled for eviction from the Cache.
type EvictCallback func(key string, value interface{})
