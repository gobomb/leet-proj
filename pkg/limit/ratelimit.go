package limit

/*
	https://github.com/golang/go/wiki/RateLimiting
*/

import (
	"context"
	"time"
)

const rateLimit = time.Second / 10 // 10 calls per second

// Client is an interface that calls something with a payload.
type Client interface {
	Call(*Payload)
}

// Payload is some payload a Client would send in a call.
type Payload struct{}

// RateLimitCall rate limits client calls with the payloads.
func RateLimitCall(client Client, payloads []*Payload) {
	ticker := time.NewTicker(rateLimit)
	defer ticker.Stop()

	for _, payload := range payloads {
		// 以一定的频率阻塞住，1秒内只产生10个新的goroutine在处理
		<-ticker.C // rate limit our client calls
		go client.Call(payload)
	}
}

// BurstRateLimitCall allows burst rate limiting client calls with the
// payloads.
func BurstRateLimitCall(ctx context.Context, client Client, payloads []*Payload, burstLimit int) {
	throttle := make(chan time.Time, burstLimit)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		ticker := time.NewTicker(rateLimit)
		defer ticker.Stop()
		for t := range ticker.C {
			select {
			case throttle <- t:
			case <-ctx.Done():
				return // exit goroutine when surrounding function returns
			}
		}
	}()

	// 没有突发流量时，1秒内只产生10个新的goroutine在处理
	// 有突发流量，那一刻可以成为例外，处理 burstLimit 个请求，之前和之后还是1秒10个的速率
	for _, payload := range payloads {
		<-throttle // rate limit our client calls
		go client.Call(payload)
	}
}
