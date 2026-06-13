package main

import (
	"fmt"
	"rateLimiter/pkg"
	"sync"
	"time"
)

func main() {
	cfg := pkg.Config{
		Capacity:    5,
		RefreshRate: 1,
	}

	ctrl := pkg.NewController(pkg.TOKENBUCKET, &cfg)
	if ctrl == nil {
		return
	}
	defer ctrl.Shutdown()

	// Send burst count 10 requests at once
	sendBurst(ctrl, 10, "")

	// wait for 5 seconds to refill
	fmt.Println()
	fmt.Println("==========5 seconds pause==========")
	fmt.Println()
	time.Sleep(5 * time.Second)

	// Again send burst of 6 requests at once
	sendBurst(ctrl, 6, "")

	// burst for ratelimit keys (for specific keys)
	for _, user := range []string{"user1", "user2", "user3"} {
		fmt.Printf("\nRequests for %s:\n", user)
		sendBurst(ctrl, 7, user)
	}

	// wait for 5 seconds to refill
	fmt.Println()
	fmt.Println("==========5 seconds pause==========")
	fmt.Println()
	time.Sleep(5 * time.Second)

	// High concurrency scenario
	count := 20
	results := make([]<-chan bool, 0, count)
	for i := 0; i < count; i++ {
		results = append(results, ctrl.AllowAsync(""))
	}

	allowed := 0
	for _, ch := range results {
		if <-ch {
			allowed++
		}
	}
	fmt.Printf("Total request: [%v], Passed: [%v], Failed: [%v]\n", count, allowed, count-allowed)
}

func sendBurst(ctrl *pkg.Controller, count int, rateLimitKey string) {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		allowed int
	)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if ctrl.Allow(rateLimitKey) {
				mu.Lock()
				allowed++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Total requestes: [%v], Passed: [%v], Failed: [%v]\n", count, allowed, count-allowed)
}
