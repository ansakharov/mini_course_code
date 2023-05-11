package src

import (
	"io"
)

type Reader interface {
	Read(p []byte) (int, error)
	ReadAll(bufSize int) (string, error)
	BytesRead() int64
}

type CountingReaderImplementation struct {
	Reader         io.Reader
	TotalBytesRead int64
}

func (cr *CountingReaderImplementation) Read(p []byte) (int, error) {
	return 0, nil
}

func NewCountingReader(r io.Reader) *CountingReaderImplementation {
	return &CountingReaderImplementation{
		Reader: r,
	}
}
