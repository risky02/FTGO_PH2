package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Product struct {
	Name string `json:"name"`
	Price int `json:"harga"`
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("content-type", "application/json")
	// 	json.NewEncoder(w).Encode(map[string]string{
	// 		"message": "hello World",
	// 	})
	// })
	
	// mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
	// 	products := []Product{
	// 		{Name: "Coca cola", Price: 10},
	// 		{Name: "Fanta", Price: 11},
	// 		{Name: "Sprite", Price: 12},
	// 	}
	// 	w.Header().Set("content-type", "application/json")
	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"message": "Hello World ini halaman products",
	// 		"products": products,
	// 	})
	// })

	router := httprouter.New()
	
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello World1",
		})
	})

	router.GET("/products", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		products := []Product{
			{Name: "Cola Coca", Price: 10},
			{Name: "Fanta", Price: 11},
			{Name: "Sprite", Price: 12},
		}
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Hello World ini halaman products",
			"products": products,
		})
	})

	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		products := []Product{
			{Name: "Cola Coca", Price: 10},
			{Name: "Fanta", Price: 11},
			{Name: "Sprite", Price: 12},
		}

		paramId := p.ByName("id")
		paramIdInt, _ := strconv.Atoi(paramId)
		product := products[paramIdInt - 1]
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Hello World ini halaman products details",
			"product_id": paramId,
			"product": product,
		})
	})

	router.POST("/products", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		newProduct := Product{}
		decoder.Decode(&newProduct)

		// validasi
		// insert db

		// response
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Hello World ini halaman create products details",
			"new_product": newProduct,
		})
	})

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "something went wrong, please try again later..",
			"detail": i,
		})
	}
	
	apps := http.Server {
		Addr: "localhost:8000",
		Handler: router,
	}
	
	err := apps.ListenAndServe()
	if err != nil {
		log.Fatalf("Server not connected")
	}
}