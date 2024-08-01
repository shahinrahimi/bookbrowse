package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

type Visitor struct {
	lastSeen time.Time
	count    int
}

type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		limit:    limit,
		window:   window,
	}

	go rl.cleanupVisitors()

	return rl
}

func (rl *RateLimiter) getVisitor(ip string) *Visitor {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	visitor, exists := rl.visitors[ip]
	if !exists || time.Since(visitor.lastSeen) > rl.window {
		rl.visitors[ip] = &Visitor{lastSeen: time.Now()}
		visitor = rl.visitors[ip]
	}

	return visitor

}

func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, visitor := range rl.visitors {
			if time.Since(visitor.lastSeen) > rl.window {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		visitor := rl.getVisitor(ip)
		visitor.lastSeen = time.Now()

		if visitor.count > rl.limit {
			utils.WriteJSON(rw, http.StatusTooManyRequests, types.ApiError{Error: "rate limit exeeded"})
			return
		}

		visitor.count++
		next.ServeHTTP(rw, r)
	})
}
