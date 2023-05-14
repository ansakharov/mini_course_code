package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const (
	port = 8000
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customValue := r.Header.Get("X-Custom-Value")

		// Store the custom value in the context
		ctx := context.WithValue(r.Context(), "id", customValue)
		ctx = context.WithValue(ctx, "message", r.URL.Query().Get("message"))

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Printf("X-Custom-Value ID %s: finished processing\n", customValue)
	})
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	response := businessFunc(r.Context(), message)

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(response))
}

func businessFunc(ctx context.Context, msg string) string {
	value := ctx.Value("id")

	fmt.Printf("Let's log id: %v\n", value)

	return fmt.Sprintf("Echo: %s\n", msg)

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", middleware(http.HandlerFunc(echoHandler)))

	fmt.Printf("Server running at %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
