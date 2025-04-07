package main

import (
	"example/hello/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	log.Println("Server running at port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
