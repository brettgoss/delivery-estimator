package main

import (
	"fmt"
	"log"
	"net/http"
)

func deliveryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/deliveries" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		log.Println("/deliveries: POST received")

		err := r.ParseForm()
		if err != nil {
			log.Print(err)
		}
		log.Printf("/deliveries: Form data: %+v\n", r.Form)
		fmt.Fprintf(w, "Delivery request received: %+v\n", r.Form)
		return
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "Delivery request received!")
		return
	}

	http.Error(w, "Method is not supported.", http.StatusNotFound)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/deliveries", deliveryHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
