package cache

import (
	"time"
)

type CacheManager struct {
	redisClient *RedisClient
	localCache  *LRUCache
}

func NewCacheManager(redisClient *RedisClient, localCache *LRUCache) *CacheManager {
	return &CacheManager{
		redisClient: redisClient,
		localCache:  localCache,
	}
}

func (cm *CacheManager) Get(key string) (map[string]any, error) {
	// Try local cache
	if val, ok := cm.localCache.Get(key); ok {
		return val.(map[string]any), nil
	}

	// Try Redis
	val, err := cm.redisClient.Get(key)
	if err == nil {
		// Set to local cache for faster next read
		cm.localCache.Set(key, val)
	}
	return val.(map[string]any), err
}

func (cm *CacheManager) Set(key string, value map[string]any, expiration time.Duration) error {
	cm.localCache.Set(key, value)
	return cm.redisClient.Set(key, value, expiration)
}

func (cm *CacheManager) Delete(key string) error {
	cm.localCache.Delete(key)
	return cm.redisClient.Delete(key)
}
