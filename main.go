package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received...")

	time.Sleep(2 * time.Second)
	fmt.Fprintln(w, "Yo!")

	log.Println("Request served.")
}
