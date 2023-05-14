package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	serverURL = "http://localhost:8000"
)

func echoRequest(message string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	q := req.URL.Query()
	q.Add("message", message)
	req.URL.RawQuery = q.Encode()

	startTime := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf(
			"Request took longer than 1.5 second, canceling the request. (Elapsed time: %.2f seconds), err: %w\n",
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

func main() {
	message := "Hello, echo server!"
	if err := echoRequest(message); err != nil {
		log.Fatal(err)
	}
}
