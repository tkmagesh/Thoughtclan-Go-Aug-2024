package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
}

func NewAppServer() *AppServer {
	return &AppServer{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (appServer *AppServer) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.handlers[pattern] = handler
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\n", r.Method, r.URL.Path)
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

func main() {
	appServer := NewAppServer()
	appServer.Register("/", IndexHandler)
	appServer.Register("/products", ProductsHandler)
	appServer.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
