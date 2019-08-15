package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	DOB  string `json:"dob"`
	Home string `json:"home"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://faizan_admin0:<password>@faizan-cluster0-iejkd.mongodb.net/test?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("users_db").Collection("ceb")

	// goodMan := user{
	// 	Name: "faizan", Age: 23, DOB: "haha", Home: "har jgah",
	// }
	//result, _ := collection.InsertOne(ctx, goodMan)

	var goodMan user
	// id, _ := primitive.ObjectIDFromHex("5d555c2aa7cd94f7e5812922")
	// result := collection.FindOne(context.Background(), bson.M{"_id": id})
	// result.Decode(goodMan)

	objID := "5d555c2aa7cd94f7e5812922"
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&goodMan)

	//fmt.Print(result.InsertedID)
	fmt.Println(goodMan)
}
