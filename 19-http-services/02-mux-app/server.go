package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	time.Sleep(10 * time.Second)
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
	// id := r.PathValue("product_id")
	vars := mux.Vars(r)
	id := vars["id"]
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
func logMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\n", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func profileMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		elapsed := time.Since(start)
		fmt.Println("Elapsed : ", elapsed)
	})
}

func timeoutMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		newReq := r.WithContext(timeoutCtx)
		handler.ServeHTTP(w, newReq)
		if newReq.Context().Err() == context.DeadlineExceeded {
			log.Println("timeout occurred!")
			w.WriteHeader(http.StatusRequestTimeout)
		}

	})
}

var authKeyUserMapping map[string]string = map[string]string{
	"key-1": "user-1",
	"key-2": "user-2",
	"key-3": "user-3",
	"key-4": "user-4",
}

func authorizationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authKey, exists := r.Header["Authorization"]; exists {
			/*
				Get the userid for the authkey from authKeyUserMapping and make it accessible in the handler
			*/
			if userid, exists := authKeyUserMapping[authKey[0]]; exists {
				valCtx := context.WithValue(r.Context(), "user-id", userid)
				handler.ServeHTTP(w, r.WithContext(valCtx))
				return
			}
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
	})
}

func main() {
	fmt.Println("Process ID :", os.Getpid())
	router := mux.NewRouter()
	// middlewares
	router.Use(logMiddleware)
	// router.Use(timeoutMiddleware)
	// router.Use(authorizationMiddleware)
	router.Use(profileMiddleware)

	// handlers
	router.Handle("/", authorizationMiddleware(http.HandlerFunc(IndexHandler))).Methods(http.MethodGet)
	router.HandleFunc("/products", GetProductsHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", NewProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/products/{id:[0-9]+}", GetOneProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", CustomersHandler).Methods(http.MethodGet)

	/*
		http.Handle("/", router)
		http.ListenAndServe(":8080", nil)
	*/

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	log.Println("Received shutdown signal!")
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
