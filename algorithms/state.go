package main

type rateLimterState struct {
	tokenBucket *tokenBucket
	slidingWindow *slidingWindow
	EWMA *EWMA
}