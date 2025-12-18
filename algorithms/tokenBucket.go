package main

import "time"

type tokenBucket struct {
	Capacity float64
	RefillRate float64
	Tokens float64
	LastRefill float64
}

func (tb *tokenBucket) Allow(n float64) bool{
	now := time.Now().Unix()

	elapsed := float(now - tb.LastRefill)

	if elapsed > 0 {
		tb.Tokens += elapsed * tb.RefillRate

		if tb.Tokens > tb.Capacity {
			tb.Tokens = tb.Capacity
		}
		tb.LastRefill = now
	}

	if tb.Tokens >= n {
		tb.Tokens -= n
		return true
	}

	return false
}