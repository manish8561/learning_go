package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/single/examples"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
* There are 3 types of design patterns:
* Creational -> Singleton, Factory Method, Abstract Factory
* Structural -> Adaptor, Bridge, etc.
* Behavioural -> Command, Interperter, etc.
 */
// one of the design pattern in golang Creational Design pattern
// singleton desing pattern below code
var lock = &sync.Mutex{}

type Single struct {
	connection *mongo.Client
}

var signletonInstance *Single

func getSingleTonInstance() *Single {
	lock.Lock()
	defer lock.Unlock()

	if signletonInstance == nil {
		fmt.Println("Creating DB connection")
		uri := "localhost:27017/testingDB"

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
		if err != nil {
			fmt.Println(err)
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Println(err)
		}

		signletonInstance = &Single{connection: client}
	} else {
		fmt.Println("DB connection is already created!!!")
	}
	return signletonInstance
}

func main() {
	// another example of singleton
	examples.Example1()

	
	for i := 0; i < 30; i++ {
		go getSingleTonInstance()
	}
	fmt.Scanln()

}
