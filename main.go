package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Answer struct {
	ID       string `json:"id,omitempty"`
	Answer_1 string `json:"answer1,omitempty"`
	Answer_2 string `json:"answer2,omitempty"`
	Answer_3 string `json:"answer3,omitempty"`
}

var client *mongo.Client

func main() {
	// Set up MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	router.HandleFunc("/", GetQuestion).Methods("GET")
	router.HandleFunc("/", SubmitAnswer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("What is your favorite programming language?")
}

func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	var answer Answer
	_ = json.NewDecoder(r.Body).Decode(&answer)

	collection := client.Database("surveyDB").Collection("answers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, bson.M{"Answer1": answer.Answer_1, "Answer2": answer.Answer_2, "Answer3": answer.Answer_3})
	json.NewEncoder(w).Encode(result.InsertedID)
}
