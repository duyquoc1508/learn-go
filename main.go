package main

import (
	driver "go-demo/drivers"

	// models "go-demo/model"
	handler "go-demo/handlers"
	middleware "go-demo/middlewares"
	"log"
	"net/http"
	// repoImpl "go-demo/repository/repoImpl"
)

func main() {
	driver.ConnectMongoDB()

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	// converting our handler function to handler type to make use of our middleware
	http.Handle("/user", middleware.Auth(http.HandlerFunc(handler.GetUser)))

	log.Println("Server running at [:8080]")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
