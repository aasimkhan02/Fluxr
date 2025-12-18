package main

type rateLimterState struct {
	TokenBucket *tokenBucket
	SlidingWindow *slidingWindow
	Ewma *ewma
}