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
	getRouter.HandleFunc("/", np.GetProducts)

	s := &http.Server{ //creating a server to tune the elements to timeout to avoid bad connections errors from server
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
