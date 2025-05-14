package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)	
	log.Println("Received request for:", r.URL.Path)
	w.Write([]byte("Hello, World!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)	
	log.Println("Received request for:", r.URL.Path)
	w.Write([]byte("Snippet View"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)	
	log.Println("Received request for:", r.URL.Path)
	w.Write([]byte("Snippet Create"))
}

func handleProductsItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)	
	log.Println("Received request for:", r.URL.Path)
	log.Println("Product Item ID:", r.PathValue("itemID"))
	w.Write([]byte("Product Item"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/products/item/{itemID}", handleProductsItem)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}	
}