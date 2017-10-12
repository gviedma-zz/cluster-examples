package shared

import (
	"fmt"
	"log"

	actor "github.com/AsynkronIT/protoactor-go/actor"
)

type HelloCountActor struct {
	count int
}

func (a *HelloCountActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Printf("Started HelloCountActor %v\n", ctx.Self().Id)
	case *actor.Stopped:
		fmt.Printf("Stopped HelloCountActor %v\n", ctx.Self().Id)
	case *HelloRequest:
		a.count++
		m := fmt.Sprintf("Hello %s from %s, my count is %d", msg.Name, ctx.Self().Id, a.count)
		ctx.Respond(&HelloResponse{Message: m})
	default:
		log.Printf("Unknown message %v", msg)
	}
}
