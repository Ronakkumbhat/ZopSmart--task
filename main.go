package main

import (
	"strconv"
    "encoding/json"
    "fmt"
    "gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()
	var books []string

	// register route greet
	app.GET("/addbook", func(ctx *gofr.Context) (interface{}, error) {
		// Get data from request body
		var requestData struct {
			Name string `json:"Name"`
		}

		// Use json.NewDecoder to decode the request body into requestData
		err := json.NewDecoder(ctx.Request().Body).Decode(&requestData)
		if err != nil {
			return nil, err
		}

		// Insert the value into the array
		books = append(books, requestData.Name)
		// Access the data
		fmt.Println("Hello", books)

		// Get the last element of the slice
		lastBook := books[len(books)-1]

		return "Hello " + lastBook,nil
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
	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}