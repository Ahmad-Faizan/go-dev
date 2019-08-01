package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://faizan_admin0:<password>@faizan-cluster0-iejkd.mongodb.net/test?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("This is the client info")
	fmt.Println(client)
	fmt.Println("Connected to MongoDB cluster")

	//Lists the Databases residing in the cluster
	//str, err := client.ListDatabaseNames(ctx, bson.M{})
	//fmt.Println(str)

}
