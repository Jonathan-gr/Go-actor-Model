package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type Inventory struct {
	Bottles int
}
type myEvent struct {
	Info string
}
type Player struct {
	HP           int
	inverntoryID *actor.PID
}

func newInventory(bottles int) actor.Producer {
	return func() actor.Receiver {
		return &Inventory{Bottles: bottles}
	}
}
func newPlayer(hp int) actor.Producer {
	return func() actor.Receiver {
		return &Player{HP: hp}
	}
}

type drinkBottle struct {
	Amount int
}

func (i *Inventory) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("Inventory started with bottles:", i.Bottles)
		c.Engine().Subscribe(c.PID())
	case actor.Stopped:
		fmt.Println("Inventory stopped")
	case drinkBottle:
		fmt.Println("Drinking bottle, healing for", msg.Amount)
	case myEvent:
		fmt.Println("Inventory received event:", msg.Info)
	}
}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("Player started with HP:", p.HP)
		// Spawn an inventory actor as a child
		p.inverntoryID = c.SpawnChild(newInventory(3), "inventory")
		c.Engine().Subscribe(c.PID())
	case actor.Stopped:
		fmt.Println("Player stopped")
	case drinkBottle:
		c.Forward(p.inverntoryID)
	case myEvent:
		fmt.Println("Received event message", msg.Info)
	}
}
func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := e.Spawn(newPlayer(100), "player1", actor.WithID("actor12323"))

	msg := drinkBottle{Amount: 10}
	e.Send(pid, msg)
	time.Sleep(1 * time.Second)
	e.BroadcastEvent(myEvent{Info: "Game Started"})
	time.Sleep(1 * time.Second)

}
