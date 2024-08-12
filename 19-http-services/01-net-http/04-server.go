package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

/*
type AppServer struct {
	handlers    map[string]func(http.ResponseWriter, *http.Request)
	middlewares []func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)
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
*/

// Application handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// simulating a time consuming operation
LOOP:
	for range 1 {
		select {
		case <-r.Context().Done():
			break LOOP
		default:
			time.Sleep(1 * time.Second)
		}
	}
	if r.Context().Err() == context.DeadlineExceeded {
		return
	}
	fmt.Fprintf(w, "Hello, %q\n", r.Context().Value("user-id"))
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "error encoding products", http.StatusInternalServerError)
	}
}

func NewProductHandler(w http.ResponseWriter, r *http.Request) {
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
}

func GetOneProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("product_id")
	fmt.Println(id)
	if pid, err := strconv.Atoi(id); err == nil {
		for _, product := range products {
			if product.Id == pid {
				if err := json.NewEncoder(w).Encode(product); err != nil {
					http.Error(w, "error encoding product", http.StatusInternalServerError)
				}
				return
			}
		}
		http.Error(w, "requested product not found", http.StatusNotFound)
	} else {
		http.Error(w, "bad request", http.StatusBadRequest)
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

func timeoutMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		newReq := r.WithContext(timeoutCtx)
		handler(w, newReq)
		if newReq.Context().Err() == context.DeadlineExceeded {
			log.Println("timeout occurred!")
			w.WriteHeader(http.StatusRequestTimeout)
		}

	}
}

var authKeyUserMapping map[string]string = map[string]string{
	"key-1": "user-1",
	"key-2": "user-2",
	"key-3": "user-3",
	"key-4": "user-4",
}

func authorizationMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if authKey, exists := r.Header["Authorization"]; exists {
			/*
				Get the userid for the authkey from authKeyUserMapping and make it accessible in the handler
			*/
			if userid, exists := authKeyUserMapping[authKey[0]]; exists {
				valCtx := context.WithValue(r.Context(), "user-id", userid)
				handler(w, r.WithContext(valCtx))
				return
			}
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
	}
}

func main() {
	serveMux := http.DefaultServeMux
	serveMux.HandleFunc("/{$}", authorizationMiddleware(timeoutMiddleware(profileMiddleware(logMiddleware(IndexHandler)))))
	serveMux.HandleFunc("GET /products", profileMiddleware(logMiddleware(GetProductsHandler)))
	serveMux.HandleFunc("GET /products/{product_id}", profileMiddleware(logMiddleware(GetOneProductHandler)))
	serveMux.HandleFunc("POST /products", profileMiddleware(logMiddleware(NewProductHandler)))
	serveMux.HandleFunc("/customers", profileMiddleware(logMiddleware(CustomersHandler)))
	http.ListenAndServe(":8080", nil)
}
