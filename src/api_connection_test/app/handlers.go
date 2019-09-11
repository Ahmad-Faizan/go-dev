package app

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getDbConnection() (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://<username>:<password>@faizan-cluster0-iejkd.mongodb.net/test?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Fatal(err)
	}

	return client, cancel
}

//GetProfile fetches profile from database
func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	client, cancel := getDbConnection()
	defer cancel()
	collection := client.Database("users_db").Collection("ceb")

	var goodMan user

	id, _ := primitive.ObjectIDFromHex(userID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&goodMan)

	//pretty print json object
	b, _ := json.MarshalIndent(goodMan, "", "    ")
	os.Stdout.Write(b)

	render.JSON(w, r, goodMan) // A chi router helper for serializing and returning json
}

//CreateProfile creates a profile in the database
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var goodMan user

	//decode the request body to json
	json.NewDecoder(r.Body).Decode(&goodMan)

	//connect to mongodb
	client, cancel := getDbConnection()
	defer cancel()
	collection := client.Database("users_db").Collection("ceb")
	result, _ := collection.InsertOne(context.TODO(), goodMan)

	//encode the object id inserted to the response in json
	render.JSON(w, r, result)
}

//UpdateProfile updates the profile from the userID
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "userID"))

	var goodMan user

	client, cancel := getDbConnection()
	defer cancel()
	collection := client.Database("users_db").Collection("ceb")
	collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&goodMan)

	//decode the request body to json
	json.NewDecoder(r.Body).Decode(&goodMan)

	//update the document
	result, _ := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": goodMan})

	//encode the object id inserted to the response in json
	render.JSON(w, r, result)
}

//DeleteProfile deletes a profile from the database
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "userID"))

	//connect to mongodb
	client, cancel := getDbConnection()
	defer cancel()
	collection := client.Database("users_db").Collection("ceb")
	result, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	//encode the result to the response in json
	render.JSON(w, r, result)
}

//GetAllProfiles fetches a list of all profiles
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	var goodMen []user

	client, cancel := getDbConnection()
	defer cancel()
	collection := client.Database("users_db").Collection("ceb")
	cursor, _ := collection.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var goodMan user
		cursor.Decode(&goodMan)
		goodMen = append(goodMen, goodMan)
	}

	render.JSON(w, r, goodMen)
}
