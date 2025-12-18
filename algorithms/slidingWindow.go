package main

import "time"

const (
	windowSize = 10
	maxRequest = 50
)

type slidingWindow struct {
	buckets []int64
	lastSecond int64
}

// constructor 
func newSlidingWindow() *slidingWindow {
	return &slidingWindow{
		buckets: make([]int64, windowSize),
	}
}

func (sw *slidingWindow) Allow() bool {
	now := time.Now().Unix()
	index := now % windowSize

	// If time moved forward, reset this bucket
	if sw.lastSecond != now {
		sw.buckets[index] = 0
		sw.lastSecond = now
	}

	sw.buckets[index]++

	var total int64
	for _, c := range sw.buckets {
		total += c
	}

	return total <= maxRequest
}