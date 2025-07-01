package decorator

import (
	"fmt"
	"sync"
	"time"
)

// ComputeFn is any function that takes a string key and returns result+error.
type ComputeFn func(key string) (interface{}, error)

// CacheDecorator adds a thread‑safe in‑memory cache around ComputeFn.
func CacheDecorator(fn ComputeFn) ComputeFn {
	var (
		mu    sync.RWMutex
		store = make(map[string]interface{})
	)

	return func(key string) (interface{}, error) {
		mu.RLock()
		if val, ok := store[key]; ok {
			mu.RUnlock()
			return val, nil
		}
		mu.RUnlock()

		result, err := fn(key)
		if err != nil {
			return nil, err
		}

		mu.Lock()
		store[key] = result
		mu.Unlock()
		return result, nil
	}
}

func expensiveCompute(key string) (interface{}, error) {
	// simulate heavy work
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("Value for %s", key), nil
}

func RunCacheDecoratorDemo() {
	cachedCompute := CacheDecorator(expensiveCompute)
	start := time.Now()
	cachedCompute("foo") // takes ~1s
	fmt.Println("First call:", time.Since(start))

	start = time.Now()
	cachedCompute("foo") // instant from cache
	fmt.Println("Second call:", time.Since(start))
}
