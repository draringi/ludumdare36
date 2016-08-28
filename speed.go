package main

type Speed uint8

const (
	STOPPED Speed = iota
	SLOW
	STANDARD
	FAST
)

func (s Speed) KilometresPerDay() float64 {
	switch s {
	case STOPPED:
		return 0
	case SLOW:
		return 10
	case STANDARD:
		return 20
	case FAST:
		return 40
	}
	return 0
}

func (s Speed) Risk() float64 {
	switch s {
	case STOPPED:
		return 0.005
	case SLOW:
		return 0.05
	case STANDARD:
		return 0.15
	case FAST:
		return 0.40
	}
	return 1
}
