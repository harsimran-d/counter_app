package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	address := ":" + port

	globalCount := GloblaCount{}

	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			globalCount.mu.Lock()
			defer globalCount.mu.Unlock()
			globalCount.value++
			w.Write(globalCount.byteValue())
		}
	})
	log.Printf("Server started on port: %v\n", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
