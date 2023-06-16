package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// case 1: Always response success
func handleHello1(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello1" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(rw, "World")
}

// case 2 response success until counter % 5 == 4
var counter int16 = 1

func handleHello2(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello2" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method is not supported", http.StatusNotFound)
		return
	}

	if counter%5 != 4 {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(rw, "World")
	}

	log.Printf("counter : %d", counter)
	counter = counter + 1
}

// case 3 always response error
func handleHello3(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello3" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method is not supported", http.StatusNotFound)
		return
	}

	time.Sleep(10 * time.Second)

	http.Error(rw, "This is error", http.StatusInternalServerError)
}

func main() {
	//case 1 - always response success
	http.HandleFunc("/hello1", handleHello1)
	//case 2 - retry 5 times
	http.HandleFunc("/hello2", handleHello2)
	//case 3 - always response error
	http.HandleFunc("/hello3", handleHello3)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
