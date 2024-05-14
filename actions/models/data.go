package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func initConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error load env file", err)
	}
	log.Print("env successfully loaded.")

}

func initDB() *mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://" +
		os.Getenv("MONGO_USERNAME") + ":" +
		os.Getenv("MONGO_PASSWORD") +
		".@cluster0.salidj6.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(os.Getenv("MONGO_DATABASE")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client.Database(os.Getenv("MONGO_DATABASE"))
}

func main() {
	initConfig()
	db := initDB()
	//InsertQuestionCategoryBase(db)
	InsertGeneralQuestionBase(db)
}
