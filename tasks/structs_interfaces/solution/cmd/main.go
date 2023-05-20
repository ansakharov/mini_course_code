package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ansakharov/mini_course_code/tasks/structs_interfaces/solution/src"
)

func main() {
	content := "This is a Sample Text to be read by the CountingToLowerReaderImpl."
	reader := bytes.NewReader([]byte(content))
	countingReader := src.NewCountingReader(reader)

	readContent, err := countingReader.ReadAll(4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(readContent)
	fmt.Printf("\nBytes read: %d\n", countingReader.BytesRead())
}
