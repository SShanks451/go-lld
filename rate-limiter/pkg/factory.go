package pkg

import (
	"errors"
	"sync"
)

type FactoryFunc func(cfg *Config) (RateLimiter, error)

var (
	registryMutex sync.RWMutex
	registry      = map[Type]FactoryFunc{
		TOKENBUCKET: func(cfg *Config) (RateLimiter, error) {
			return newTokenBucket(cfg), nil
		},
	}
)

func New(t Type, cfg *Config) (RateLimiter, error) {
	registryMutex.RLock()
	rl, ok := registry[t]
	registryMutex.RUnlock()

	if !ok {
		return nil, errors.New("This Rate Limiter Alogrithm implemetation not found")
	}

	return rl(cfg)
}

func Register(t Type, f FactoryFunc) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	registry[t] = f
}
