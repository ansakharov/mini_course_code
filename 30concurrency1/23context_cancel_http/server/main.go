package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	port = 8000
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	sleepTime := rand.Intn(4)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	fmt.Printf("received request, sleep time: %ds\n", sleepTime)
	message := r.URL.Query().Get("message")
	response := fmt.Sprintf("Echo: %s\n", message)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", echoHandler)
	fmt.Printf("Serving on port %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
