package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	serverURL = "http://localhost:8000"
)

type key int

const (
	requestIDKey key = 0
	userInfoKey  key = 1
)

type UserInfo struct {
	ID    int
	Name  string
	Email string
}

func echoRequest(req *http.Request, message string) (err error) {
	q := req.URL.Query()
	q.Add("message", message)
	req.URL.RawQuery = q.Encode()

	startTime := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf(
			"Request took longer than 1 second, canceling the request. (Elapsed time: %.2f seconds), err: %w\n",
			time.Since(startTime).Seconds(),
			err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status code: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ReadAll %w", err)
	}

	fmt.Println(string(body))

	return
}

func reqWithMiddleWare(ctx context.Context, message string) error {
	ctx, cancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		return fmt.Errorf("NewRequestWithContext: %w", err)
	}

	req.Header.Set("X-Custom-Value", strconv.FormatInt(time.Now().Unix(), 10))

	return echoRequest(req, message)
}

func main() {
	message := "Hello, how are you?"

	ctx := context.Background()
	if err := reqWithMiddleWare(ctx, message); err != nil {
		log.Fatal(err)
	}
}
