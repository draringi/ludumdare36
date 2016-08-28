package main

import (
	"github.com/jroimartin/gocui"
	"math/rand"
	"sync"
	"time"
)

type GameLoop struct {
	paused        bool
	pausedTrigger *sync.Cond
	pausedLock    sync.Locker
	running       bool
	state         *gamestate
	kill          chan bool
	hasPaused     chan bool
	gui           *gocui.Gui
}

func (g *GameLoop) TogglePause() bool {
	g.pausedLock.Lock()
	g.paused = !g.paused
	g.pausedTrigger.Broadcast()
	p := g.paused
	g.pausedLock.Unlock()
	if p {
		g.hasPaused <- p
	}
	return p
}

func (g *GameLoop) run() {
	var timer *time.Ticker
	player := WorldState.player
	g.pausedLock.Lock()
	for g.paused {
		g.pausedTrigger.Wait()
	}
	g.pausedLock.Unlock()
	timer = time.NewTicker(activeSettings.Tick.Step())
	for {
		select {
		case <-g.hasPaused:
			g.pausedLock.Lock()
			for g.paused {
				g.pausedTrigger.Wait()
			}
			g.pausedLock.Unlock()
			timer = time.NewTicker(activeSettings.Tick.Step())
		case <-timer.C:
			// Do game logic
			player.lock.Lock()
			speed := WorldState.player.speed
			WorldState.date = WorldState.date.Next()
			player.lock.Unlock()
			if speed != STOPPED {
				distanceTraveledToday := speed.KilometresPerDay() * player.attributes.movementRate
				player.lock.Lock()
				player.kilometersTravelled += distanceTraveledToday
				player.lock.Unlock()
			} else {
				WorldState.logLock.Lock()
				WorldState.log = append(WorldState.log, NewLogEntry(WorldState.date, "You didn't go anywhere..."))
				WorldState.logLock.Unlock()
			}
			risk := speed.Risk()
			roll := rand.Float64()
			if roll < risk {
				triggerBadEvent(WorldState)
			}
			// Healing, Bad Effects, and death checks
			player.DailyChecks()

			// Check if the player is dead
			if !player.character.alive {
				g.gui.Execute(triggerGameover)
			}
			// Update UI
			g.gui.Execute(updateGUI)
		case <-g.kill:
			return
		}
	}
}

func (g *GameLoop) start() {
	if !g.running {
		g.running = true
		go g.run()
	}
}

func InitiateGameLoop(state *gamestate, gui *gocui.Gui) *GameLoop {
	g := new(GameLoop)
	g.pausedLock = new(sync.Mutex)
	g.pausedTrigger = sync.NewCond(g.pausedLock)
	g.paused = true
	g.running = false
	g.state = state
	g.kill = make(chan bool, 1)
	g.hasPaused = make(chan bool, 1)
	g.gui = gui
	return g
}

func pauseToggleButton(g *gocui.Gui, v *gocui.View) error {
	if WorldState.gameLoop != nil {
		WorldState.gameLoop.TogglePause()
		updateGUI(g)
	}
	return nil
}
