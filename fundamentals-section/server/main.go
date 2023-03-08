package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person represents a person document in the database
type Person struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Fact     string             `bson:"fact" json:"fact"`
	Location string             `bson:"location" json:"location"`
}

func main() {
	// Set the MongoDB connection URI, including the username, password, and database name
	uri := os.Getenv("MONGODB_CONNECTION_STRING")

	// Set the client options
	clientOptions := options.Client().ApplyURI(uri)

	// Set the context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB instance
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Cosmos DB MongoDB instance!")

	// Define a handler function for the "/person" endpoint
	http.HandleFunc("/person/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			createPerson(w, r, client)
		// case "DELETE":
		// 	deletePerson(w, r, client)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/person/delete", func(w http.ResponseWriter, r *http.Request) {
		// switch r.Method {
		// case "DELETE":
		// 	deletePerson(w, r, client)
		// default:
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// }
		log.Print("deletePerson called")
		deletePerson(w, r, client)
	})

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// createPerson creates a new person document in the database
func createPerson(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	// Parse the request body into a Person struct
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID field to a new unique ID
	person.ID = primitive.NewObjectID()

	// Get the MongoDB collection from the client
	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))

	// Insert the new person document into the collection
	_, err = collection.InsertOne(context.Background(), &person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response status code to 201 Created and return the newly created person's ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{ ID string }{person.ID.Hex()})
}

// deletePerson deletes a person document from the database
func deletePerson(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	// Read the request body
	decoder := json.NewDecoder(r.Body)
	var data struct {
		ID string `json:"id"`
	}
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the ID parameter
	if data.ID == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Get the MongoDB collection from the client
	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))

	// Delete the person document with the specified ID
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Person document not found", http.StatusNotFound)
		return
	}

	// Set the response status code to 204 No Content to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}
