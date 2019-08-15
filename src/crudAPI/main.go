package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-chi/chi"
)

type Profile struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	name   string             `json:"name" bson:"name"`
	age    uint8              `json:"age" bson:"age"`
	status string             `json:"status" bson:"status"`
}

var Mongodb_client *mongo.Client

func findPersonEndPoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("In find.go")
	response.Header().Set("content-type", "application/json")
	urlParams := chi.URLParam(request, "id")
	id, _ := primitive.ObjectIDFromHex(urlParams)

	var person Profile

	collection := Mongodb_client.Database("test_db").Collection("users")
	//ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)

	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	//response.Write([]byte(`{ "message": "` + err.Error + `" }`))
	// 	return
	// }
	//json.NewEncoder(response).Encode(person)
	render.JSON(response, request, person)
	fmt.Print(person)
	fmt.Print(response)
}

func main() {
	fmt.Println("Server has started....")
	//create the mongoDb connection
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	Mongodb_client, _ = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://faizan_admin0:<password>@faizan-cluster0-iejkd.mongodb.net/test?retryWrites=true&w=majority",
	))

	//create the router from chi
	router := chi.NewRouter()

	//define handlers for HTTP verbs
	router.Get("/ping", check)
	router.Get("/person/{id}", findPersonEndPoint)

	//start the api server and listen for requests
	log.Fatalln(http.ListenAndServe(":3000", router))
}

func check(writer http.ResponseWriter, router *http.Request) {
	//write a simple response
	writer.Write([]byte("welcome"))

	//write a json
	// profile := Profile{"Alex", 21, "single"}
	// js, err := json.Marshal(profile)
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// writer.Header().Set("Content-Type", "application/json")
	// writer.Write(js)
}
