package pkg

import "time"

type Type string

const (
	TOKENBUCKET   Type = "token_bucket"
	FIXEDWINDOW   Type = "fixed_window"
	SLIDINGWINDOW Type = "sliding_window"
	LEAKYBUCKET   Type = "leaky_bucket"
)

type Config struct {
	Capacity       int
	RefreshRate    int
	RefillInterval time.Duration
}

func (c *Config) withDefaults() *Config {
	if c.Capacity <= 0 {
		c.Capacity = 10
	}
	if c.RefreshRate <= 0 {
		c.RefreshRate = 10
	}
	if c.RefillInterval <= 0 {
		c.RefillInterval = time.Second
	}

	return c
}
