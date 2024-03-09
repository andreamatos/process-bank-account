package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Account struct represents the structure of each account in the JSON file
type Account struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
}

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Read input JSON file
	inputFile := "input.json"
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error reading input file:", err)
	}

	// Unmarshal JSON data into a slice of Account structs
	var accounts []Account
	if err := json.Unmarshal(inputData, &accounts); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	// Access the "accounts" collection in the "bank" database
	collection := client.Database("bank").Collection("accounts")

	// Process account data (e.g., update balances or statuses)
	for _, acc := range accounts {
		// Example: Increase balance by 10%
		acc.Balance *= 1.1

		// Example: Change status to "Active" if balance is positive, "Inactive" otherwise
		if acc.Balance > 0 {
			acc.Status = "Active"
		} else {
			acc.Status = "Inactive"
		}

		// Insert or update the account data in the MongoDB collection
		filter := bson.M{"id": acc.ID}
		update := bson.M{"$set": acc}
		_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
		if err != nil {
			log.Fatal("Error updating MongoDB:", err)
		}
	}

	fmt.Println("Processing complete. Data updated in MongoDB.")
}
