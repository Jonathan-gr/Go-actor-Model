package main

import (
	"log"

	"github.com/anthdm/hollywood/actor"
)

type Player struct {
	HP int
}

func newPlayer() actor.Receiver {
	return &Player{}
}

func (p *Player) Receive(c *actor.Context) {}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatalf("failed to create actor engine: %v", err)
	}
	e.Spawn(newPlayer, "player")
}
