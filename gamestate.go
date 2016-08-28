package main

import "sync"

type gamestate struct {
	player   *Player
	location coords
	date     Date
	logLock  sync.Locker
	log      []LogEntry
	gameLoop *GameLoop
}

var WorldState *gamestate
