package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/cluster/consul"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gviedma/cluster-examples/shared"
)

const (
	timeout = 1 * time.Second
)

func main() {
	remote.Register("Hello", actor.FromProducer(func() actor.Actor {
		return &shared.HelloCountActor{}
	}))

	cp, err := consul.New()
	if err != nil {
		log.Fatal(err)
	}
	cluster.Start("mycluster", "cluster-example-member:8081", cp)

	inc()

	log.Println("Sleeping")
	time.Sleep(1 * time.Hour)
}

func inc() {
	time.Sleep(15 * time.Second)
	fmt.Println("Done sleeping")
	for j := 0; j < 5; j++ {

		for i := 0; i < 20; i++ {
			name := fmt.Sprintf("hello%d", i)
			helloPid, statusCode := cluster.Get(name, "Hello")
			fmt.Printf("Making request to %v (%v)\n", helloPid, statusCode)
			if helloPid != nil {
				result, err := helloPid.RequestFuture(&shared.HelloRequest{Name: fmt.Sprintf("loop-%d", i)}, 15*time.Second).Result()
				fmt.Printf("Got response %v (%v)\n", result, err)
			} else {
				fmt.Printf("Got error resolving %v: %v\n", name, statusCode)
			}
		}
	}
}
