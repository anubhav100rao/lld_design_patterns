package constructor

import (
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
}

type memoryCache struct {
	store map[string]string
	group singleflight.Group
}

func (m *memoryCache) Get(key string) (string, error) {
	if value, exists := m.store[key]; exists {
		return value, nil
	}
	return "", nil // or an error if preferred
}
func (m *memoryCache) Set(key string, value string, ttl time.Duration) error {
	m.store[key] = value
	if ttl > 0 {
		go func() {
			time.Sleep(ttl)
			delete(m.store, key)
		}()
	}
	return nil
}

type redisCache struct {
	client *redis.Client
}

func (r *redisCache) Get(key string) (string, error) {
	val, err := r.client.Get(r.client.Context(), key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err // Other error
	}
	return val, nil
}
func (r *redisCache) Set(key string, value string, ttl time.Duration) error {
	if ttl > 0 {
		return r.client.Set(r.client.Context(), key, value, ttl).Err()
	}
	return r.client.Set(r.client.Context(), key, value, 0).Err() // No expiration
}

type CacheConfig struct {
	Type     string // "memory" or "redis"
	RedisOpt *redis.Options
}

func NewCache(config CacheConfig) (Cache, error) {
	switch config.Type {
	case "memory":
		return &memoryCache{store: make(map[string]string)}, nil
	case "redis":
		if config.RedisOpt == nil {
			return nil, nil // or an error if preferred
		}
		client := redis.NewClient(config.RedisOpt)
		if err := client.Ping(client.Context()).Err(); err != nil {
			return nil, err
		}
		return &redisCache{client: client}, nil
	default:
		return nil, nil // or an error if preferred
	}

}
