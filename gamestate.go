package main

import "sync"

type gamestate struct {
	player       *Player
	nextLocation int
	date         Date
	logLock      sync.Locker
	log          []LogEntry
	gameLoop     *GameLoop
}

var WorldState *gamestate
