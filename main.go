package main

import (
	"log"

	"github.com/anthdm/hollywood/actor"
)

type Player struct {
	HP int
}

func newPlayer(hp int) actor.Producer {
	return func() actor.Receiver {
		return &Player{HP: hp}
	}
}

func (p *Player) Receive(ctx *actor.Context) {
}
func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	e.Spawn(newPlayer(100), "player1")
}
