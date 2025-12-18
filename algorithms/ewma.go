package main

type EWMA struct {
	alpha float64
	value float64
	prev float64
	initialized bool
}

func NewEWMA(alpha float64) *EWMA {
	return &EWMA{
		alpha: alpha,
	}
}

func (e *EWMA) Update(current_rate float64) {
	if !e.initialized {
		e.value = currentRate
		e.prev = currentRate
		e.initialized = true
		return
	}

	e.prev = e.value
	e.value = e.alpha * currentRate + (1.0 - e.alpha) * e.value
}

func (e *EWMA) Derivative() float64 {
	if !e.initialized {
		return 0
	}
	return e.value - e.prev
}

func (e *EWMA) Value() float64 {
	return e.value
}