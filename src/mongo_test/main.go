package main

import (
	"context"
	"fmt"
	"log"
	"strings"
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

func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://faizan_admin0:<password>@faizan-cluster0-iejkd.mongodb.net/test?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getChoice() string {
	var choice string

	fmt.Println("Database Options: \n Insert \n Update \n Delete \n Find")
	fmt.Print("Enter the choice: ")
	fmt.Scan(&choice)

	return strings.ToLower(choice)
}

func dbOperation(choice string) {
	client := connect()
	collection := client.Database("users_db").Collection("ceb")

	switch choice {
	case "insert":
		insertInDB(client, collection)
	case "update":
		updateDB(client, collection)
	case "delete":
		deleteFromDB(client, collection)
	case "find":
		searchDB(client, collection)
	default:
		fmt.Println("Invalid Option entered")
	}
}

func insertInDB(client *mongo.Client, collection *mongo.Collection) {
	person := createPerson()
	result, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		fmt.Println("error in insertion", err)
	}
	fmt.Println("Insert successful with id:", result.InsertedID)
}

func updateDB(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user
	fmt.Print("Enter the id of the person: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)
	person = updateField(person)
	collection.UpdateOne(context.TODO(), bson.M{"_id": id}, person)
	fmt.Println(person, "person updated successfully")
}

func deleteFromDB(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user
	fmt.Print("Enter the id of the person: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)
	fmt.Println(person, "person deleted successfully")
}

func searchDB(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user
	fmt.Print("Enter the id of the person: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)
	if person.Age != 0 {
		fmt.Print(person, " found successfully")
	} else {
		fmt.Print("Search failed")
	}
}

func createPerson() interface{} {
	var name, dob, home string
	var age int

	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter age: ")
	fmt.Scan(&age)
	fmt.Print("Enter date of birth (dd-mm-yyyy format): ")
	fmt.Scan(&dob)
	fmt.Print("Enter your hometown: ")
	fmt.Scan(&home)

	person := user{Name: name, Age: age, DOB: dob, Home: home}

	return person
}

func updateField(person user) user {
	var name, dob, home string
	var age int
	var field string

	fmt.Println("Enter the field you want to update: \n Name \n Age \n dob \n home")
	fmt.Print("Enter one of the above fields: ")
	fmt.Scan(&field)

	switch field {
	case "name":
		fmt.Print("Enter new name: ")
		fmt.Scan(&name)
		person.Name = name
	case "age":
		fmt.Print("Enter new age: ")
		fmt.Scan(&age)
		person.Age = age
	case "dob":
		fmt.Print("Enter new dob (in dd-mm-yyyy format): ")
		fmt.Scan(&dob)
		person.DOB = dob
	case "home":
		fmt.Print("Enter new home: ")
		fmt.Scan(&home)
		person.Home = home
	}
	return person
}

func main() {
	dbOperation(getChoice())

}
