package util

import (
	"time"

	"golang.org/x/time/rate"
	"k8s.io/client-go/util/workqueue"
)

type RatelimitConfig struct {
	QueueTokenRefillRate int
	QueueTokenBucketSize int
	MaxDelay             time.Duration
}

func NewRateLimiter(config RatelimitConfig) workqueue.RateLimiter {
	ratelimiter := workqueue.NewWithMaxWaitRateLimiter(
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(config.QueueTokenRefillRate), config.QueueTokenBucketSize)},
		config.MaxDelay,
	)
	return ratelimiter
}

func CreateNamedRateLimitingQueue(name string, config RatelimitConfig) workqueue.RateLimitingInterface {
	rateLimiter := NewRateLimiter(config)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimiter, name)
	return queue
}
