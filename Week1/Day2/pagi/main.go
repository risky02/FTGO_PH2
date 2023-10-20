package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	Name string `json:"name"`
	Price int `json:"harga"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello World",
		})
	})
	
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := []Product{
			{Name: "Coca cola", Price: 10},
			{Name: "Fanta", Price: 11},
			{Name: "Sprite", Price: 12},
		}
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Hello World ini halaman products",
			"products": products,
		})
	})

	apps := http.Server {
		Addr: "localhost:8000",
		Handler: mux,
	}
	
	err := apps.ListenAndServe()
	if err != nil {
		log.Fatalf("Server not connected")
	}
}