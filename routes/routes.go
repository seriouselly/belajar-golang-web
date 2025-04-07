package routes

import (
	"example/hello/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/product", handlers.ProductHandler)
}
