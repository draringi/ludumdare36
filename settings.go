package main

import (
	"time"
)

var activeSettings Settings

type GameSpeed uint8

const (
	gameSpeedSlow GameSpeed = iota
	gameSpeedMedium
	gameSpeedFast
)

func (g GameSpeed) Step() time.Duration {
	switch g {
	case gameSpeedFast:
		return time.Second
	case gameSpeedMedium:
		return 2 * time.Second
	case gameSpeedSlow:
		return 4 * time.Second
	}
	return time.Hour
}

type Settings struct {
	Tick GameSpeed
}
