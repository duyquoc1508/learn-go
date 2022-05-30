package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	model "restapi.com/main/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func homeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome home!")
}

func createPerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person model.Person
	fmt.Println(request.Body)
	json.NewDecoder(request.Body).Decode(&person)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(person)
	result, _ := collection.InsertOne(ctx, person)
	id := result.InsertedID
	person.ID = id.(primitive.ObjectID)
	json.NewEncoder(response).Encode(person)
}

func getPersonByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person model.Person
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println(id)
	defer cancel()
	err := collection.FindOne(ctx, model.Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func main() {
	if err := godotenv.Load(); err != nil { // err is scoped only to the if block that handle it
		log.Fatal("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variables")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Set client option
	clientOptions := options.Client().ApplyURI(uri)
	// Connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("example").Collection("people")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/people", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", getPersonByID).Methods("GET")

	log.Printf("Server running at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
