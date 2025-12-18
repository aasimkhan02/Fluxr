package main

const (
	ewmaBurstThreshold = 5.0
)

func (state *rateLimterState) Admit(tokensRequested float64, currentRate float64) bool {
	if state.tokenBucket.Allow(tokensRequested){
		return true
	}

	if state.slidingWindow(){
		return true
	}

	state.EWMA.Update(currentRate)

	if state.EWMA.Derivative() > ewmaBurstThreshold {
		return true
	}

	return false
}