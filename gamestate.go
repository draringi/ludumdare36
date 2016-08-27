package main

type gamestate struct {
  player *Player
  location coords
}

var WorldState *gamestate
