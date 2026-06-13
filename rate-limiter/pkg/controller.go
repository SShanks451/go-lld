package pkg

import (
	"fmt"
)

type Controller struct {
	rl RateLimiter
}

func NewController(t Type, cfg *Config) *Controller {
	rl, err := New(t, cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &Controller{
		rl: rl,
	}
}

func (c *Controller) Allow(key string) bool {
	isAllowed := c.rl.Allow(key)
	display := "global"
	if key != "" {
		display = key
	}

	if isAllowed {
		fmt.Printf("Request with key [%v] allowed\n", display)
	} else {
		fmt.Printf("Request with key [%v] blocked\n", display)
	}

	return isAllowed
}

func (c *Controller) AllowAsync(key string) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		ch <- c.Allow(key)
	}()
	return ch
}

func (c *Controller) UpdateConfig(cfg Config) {
	c.rl.UpdateConfig(cfg)
}

func (c *Controller) Shutdown() {
	c.rl.Shutdown()
}
