package pkg

import (
	"sync"
	"sync/atomic"
	"time"
)

type TokenBucket struct {
	capacity    int
	refreshRate atomic.Int64
	global      *Bucket
	mu          sync.RWMutex
	buckets     map[string]*Bucket
	done        chan struct{}
	stopOnce    sync.Once
}

func newTokenBucket(cfg *Config) *TokenBucket {
	cfg = cfg.withDefaults()

	tb := TokenBucket{}
	tb.capacity = cfg.Capacity
	tb.global = newBucket(cfg.Capacity)
	tb.buckets = make(map[string]*Bucket)
	tb.done = make(chan struct{})
	tb.refreshRate.Store(int64(cfg.RefreshRate))

	go tb.refillLoop(cfg.RefillInterval)
	return &tb
}

func (tb *TokenBucket) refillLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rate := tb.refreshRate.Load()
			tb.global.refill(int(rate), tb.capacity)

			tb.mu.RLock()
			for _, b := range tb.buckets {
				b.refill(int(rate), tb.capacity)
			}
			tb.mu.RUnlock()
		case <-tb.done:
			return
		}
	}
}

func (tb *TokenBucket) userBucket(key string) *Bucket {
	tb.mu.RLock()
	bucket, ok := tb.buckets[key]
	tb.mu.RUnlock()
	if ok {
		return bucket
	}

	tb.mu.Lock()
	defer tb.mu.Unlock()
	bucket, ok = tb.buckets[key]
	if !ok {
		bucket = newBucket(tb.capacity)
		tb.buckets[key] = bucket
	}

	return bucket
}

func (tb *TokenBucket) Allow(key string) bool {
	if key == "" {
		return tb.global.tryConsume()
	}

	b := tb.userBucket(key)
	return b.tryConsume()
}

func (tb *TokenBucket) UpdateConfig(cfg Config) {
	if cfg.RefreshRate > 0 {
		tb.refreshRate.Store(int64(cfg.RefreshRate))
	}
}

func (tb *TokenBucket) Shutdown() {
	tb.stopOnce.Do(func() {
		close(tb.done)
	})
}
