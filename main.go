package main

import (
	"fmt"
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

type takeDamage struct{}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case takeDamage:
		fmt.Println("Taking damage!", msg)
	}
}
func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := e.Spawn(newPlayer(100), "player1", actor.WithID("actor12323"))
	fmt.Println(pid)
}
