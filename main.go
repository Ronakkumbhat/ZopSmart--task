package main

import (
	"strconv"
	"strings"
    "encoding/json"
    "fmt"
    "gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
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
    collection := client.Database("test").Collection("books")

	 if err != nil {
		 log.Fatal(err)
	 }
 
	 // Check the connection
	 err = client.Ping(context.TODO(), nil)
 
	 if err != nil {
		 log.Fatal(err)
	 }
 
	 fmt.Println("Connected to MongoDB!")
 
	// register route greet
	app.GET("/addbook", func(ctx *gofr.Context) (interface{}, error) {
		var book Book
		err := json.NewDecoder(ctx.Request().Body).Decode(&book)
		if err != nil {
			return nil, err
		}
	
		// Insert the book into the collection
		insertResult, err := collection.InsertOne(context.TODO(), book)
		if err != nil {
			return nil, err
		}
	
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		return "Book inserted with ID: " + insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
	})
	app.GET("/updatestatus", func(ctx *gofr.Context) (interface{}, error) {
		// Get data from request body
		var requestData struct {
			Name string `json:"Name"`
			Status string `json:"Status"`
		}

		// Use json.NewDecoder to decode the request body into requestData
		err := json.NewDecoder(ctx.Request().Body).Decode(&requestData)
		if err != nil {
			return nil, err
		}

		var tempbooks []string
		if(requestData.Status == "returned"){
			books = append(books, requestData.Name)
		}else if(requestData.Status == "borrowed"){
			for i := 0; i < len(books); i++ {
				if(books[i] != requestData.Name){
					tempbooks = append(tempbooks, books[i])
				}
			}
			books = tempbooks
		}else if(requestData.Status == "lost"){
			for i := 0; i < len(books); i++ {
				if(books[i] != requestData.Name){
					tempbooks = append(tempbooks, books[i])
				}
			}
			books = tempbooks
		}else{
			fmt.Println("invalid", books, requestData)
		
		}
		return "Hello " + strconv.Itoa(len(books)),nil
	})
	app.GET("/getstatus", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello " + strings.Join(books, ", "), nil
	})
	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}