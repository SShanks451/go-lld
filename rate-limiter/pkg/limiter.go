package pkg

type RateLimiter interface {
	Allow(key string) bool
	UpdateConfig(cfg Config)
	Shutdown()
}
