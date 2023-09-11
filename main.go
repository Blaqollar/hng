package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Blaqollar/hng/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	np := handlers.NewProducts(l)
	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/api", np.GetProducts)

	//Setting up the server //tune the elements to timeout to avoid bad connections errors
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
