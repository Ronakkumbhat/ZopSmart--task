package main

import (
	"strings"
    "encoding/json"
    "fmt"
    "gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
    "log"
	"context"
)
func main() {
	// initialise gofr object
	app := gofr.New()
	var books []string
	type Book struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	clientOptions := options.Client().ApplyURI("mongodb+srv://ronakkumbhat8:password1234@cluster0.krmii7s.mongodb.net/")
	 // Connect to MongoDB
	 client, err := mongo.Connect(context.TODO(), clientOptions)

	 if err != nil {
		 log.Fatal(err)
	 }
 
	 // Check the connection
	 err = client.Ping(context.TODO(), nil)
 
	 if err != nil {
		 log.Fatal(err)
	 }
 
	 fmt.Println("Connected to MongoDB!")
 
	app.GET("/addbook", func(ctx *gofr.Context) (interface{}, error) {
		var book Book
		book.Status="new"
		err := json.NewDecoder(ctx.Request().Body).Decode(&book)
		if err != nil {
			return nil, err
		}
		collection := client.Database("test").Collection("books")
		// Insert the book into the collection
		insertResult, err := collection.InsertOne(context.TODO(), book)
		if err != nil {
			return nil, err
		}
	
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		return "Book inserted" , nil
	})
	app.GET("/updatestatus", func(ctx *gofr.Context) (interface{}, error) {
		collection := client.Database("test").Collection("books")
		var book Book
		err := json.NewDecoder(ctx.Request().Body).Decode(&book)
		if err != nil {
			return nil, err
		}
	
		// Create a filter to find the book by name
		filter := bson.D{{"name", book.Name}}
	
		// Create an update to set the status of the book
		update := bson.D{{"$set", bson.D{{"status", book.Status}}}}
	
		// Update the book in the collection
		updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return nil, err
		}
	
		// Return the result of the update operation
		return updateResult, nil
	})
	app.GET("/getstatus", func(ctx *gofr.Context) (interface{}, error) {
		collection := client.Database("test").Collection("books")
		var book Book
		err := json.NewDecoder(ctx.Request().Body).Decode(&book)
		if err != nil {
			return nil, err
		}

		// Create a filter to find the book by name
		filter := bson.D{{"name", book.Name}}

		// Find the book in the collection
		err = collection.FindOne(context.TODO(), filter).Decode(&book)
		if err != nil {
			return nil, err
		}

		// Return the status of the book
		return book.Status, nil
			return "Hello " + strings.Join(books, ", "), nil
	})
	app.DELETE("/deletebook", func(ctx *gofr.Context) (interface{}, error) {
		collection := client.Database("test").Collection("books")
		var book Book
		err := json.NewDecoder(ctx.Request().Body).Decode(&book)
		if err != nil {
			return nil, err
		}
	
		// Create a filter to find the book by name
		filter := bson.D{{"name", book.Name}}
	
		// Delete the book from the collection
		deleteResult, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			return nil, err
		}
	
		// Return the result of the delete operation
		return deleteResult, nil
	})
	app.Start()
}