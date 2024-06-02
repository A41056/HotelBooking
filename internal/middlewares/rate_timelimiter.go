package middlewares

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	rateLimiters = make(map[uuid.UUID]*RateLimiter)
	mu           sync.Mutex
)

func getRateLimiter(userID uuid.UUID) *RateLimiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := rateLimiters[userID]
	if !exists {
		limiter = &RateLimiter{
			limiter:  rate.NewLimiter(rate.Every(time.Second), 5), // 1 request per second with burst of 5
			lastSeen: time.Now(),
		}
		rateLimiters[userID] = limiter
	}

	limiter.lastSeen = time.Now()
	return limiter
}

func cleanupOldEntries() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for userID, limiter := range rateLimiters {
			if time.Since(limiter.lastSeen) > time.Minute*5 {
				delete(rateLimiters, userID)
			}
		}
		mu.Unlock()
	}
}

func init() {
	go cleanupOldEntries()
}
