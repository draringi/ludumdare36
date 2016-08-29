package main

import "sync"

type gamestate struct {
	Player       *Player
	NextLocation int
	Date         Date
	logLock      sync.Locker
	Log          []LogEntry
	gameLoop     *GameLoop
}

var WorldState *gamestate
