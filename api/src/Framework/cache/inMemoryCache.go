package cacheinmemory

import (
	"time"

	"github.com/patrickmn/go-cache"
)

//AddNoExpiration add cache no expiration
func AddNoExpiration(name string, obj interface{}) {
	inMemory := cache.New(5*time.Minute, 10*time.Minute)
	inMemory.Set(name, obj, cache.NoExpiration)
}

//Add a cache
func Add(name string, obj interface{}, defaultExpiration time.Duration, cleanupInterval time.Duration) {
	inMemory := cache.New(cleanupInterval, cleanupInterval)
	inMemory.Set(name, obj, cache.DefaultExpiration)
}

//GetValue cache
func GetValue(cacheName string) interface{} {
	inMemory := cache.New(5*time.Minute, 10*time.Minute)
	if x, found := inMemory.Get("foo"); found {
		return x
	}

	return nil
}

//Delete a cache if it exists
func Delete(cacheName string) {
	inMemory := cache.New(5*time.Minute, 10*time.Minute)
	inMemory.Delete(cacheName)
}
