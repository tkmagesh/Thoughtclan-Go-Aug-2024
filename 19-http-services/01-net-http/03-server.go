package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Cost  float64 `json:"cost"`
	Units int     `json:"units"`
}

var products []Product = []Product{
	{100, "Pen", 10, 20},
	{101, "Pencil", 5, 25},
	{102, "Marker", 50, 10},
	{103, "Notepad", 20, 30},
}

type AppServer struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
	// middlewares []func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func NewAppServer() *AppServer {
	return &AppServer{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (appServer *AppServer) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		handler = appServer.middlewares[i](handler)
	}
	appServer.handlers[pattern] = handler
}

func (appServer *AppServer) UseMiddleware(middleware func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := appServer.handlers[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "requested resource not found", http.StatusNotFound)
}

// Application handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "error encoding products", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		newProduct.Id = len(products) + 100
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newProduct); err != nil {
			http.Error(w, "error encoding product", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All the customers list will be served")
}

// middlewares
func logMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func profileMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("Elapsed : ", elapsed)
	}
}

func main() {
	appServer := NewAppServer()

	/*
		appServer.Register("/", profileMiddleware(logMiddleware(IndexHandler)))
		appServer.Register("/products", profileMiddleware(logMiddleware(ProductsHandler)))
		appServer.Register("/customers", profileMiddleware(logMiddleware(CustomersHandler)))
	*/
	// Registering middlewares
	appServer.UseMiddleware(profileMiddleware)
	appServer.UseMiddleware(logMiddleware)

	// Registering handlers
	appServer.Register("/", IndexHandler)
	appServer.Register("/products", ProductsHandler)
	appServer.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
