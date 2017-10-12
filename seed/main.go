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

func main() {

	remote.Register("Hello", actor.FromProducer(func() actor.Actor {
		return &shared.HelloCountActor{}
	}))

	cp, err := consul.New()
	if err != nil {
		log.Fatal(err)
	}
	cluster.Start("mycluster", "cluster-example-seed:8080", cp)

	//time.Sleep(15 * time.Second)
	fmt.Println("Done sleeping")
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("hello%d", i)
		helloPid, statusCode := cluster.Get(name, "Hello")
		fmt.Printf("Making request to %v (%v)\n", helloPid, statusCode)
		if helloPid != nil {
			result, err := helloPid.RequestFuture(&shared.HelloRequest{Name: fmt.Sprintf("loop-%d", i)}, 5*time.Second).Result()
			fmt.Printf("Got response %v (%v)\n", result, err)
		} else {
			fmt.Printf("Got error resolving %v: %v\n", name, statusCode)
		}
	}

	log.Println("Sleeping")
	time.Sleep(1 * time.Hour)
}
