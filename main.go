package main

import (
    "encoding/json"
    "fmt"
    "gofr.dev/pkg/gofr"
)

func main() {
    // initialise gofr object
    app := gofr.New()

    // register route greet
    app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
        // Get data from request body
        var requestData struct {
            Name string `json:"name"`
        }
    
        // Use json.NewDecoder to decode the request body into requestData
        err := json.NewDecoder(ctx.Request().Body).Decode(&requestData)
        if err != nil {
            return nil, err
        }
    
        // Access the data
        fmt.Println("Hello", requestData.Name)

        return "Hello " + requestData.Name
    })

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
    app.Start()
}