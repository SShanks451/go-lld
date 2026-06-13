package pkg

import "sync"

type Bucket struct {
	mu     sync.Mutex
	tokens int
}

func newBucket(tokens int) *Bucket {
	return &Bucket{
		tokens: tokens,
	}
}

func (b *Bucket) tryConsume() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.tokens > 0 {
		b.tokens--
		return true
	}
	return false
}

func (b *Bucket) refill(amount, capacity int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.tokens = min(capacity, b.tokens+amount)
}
