package src

import (
	"io"
	"strings"
	"unicode"
)

type Reader interface {
	Read(p []byte) (int, error)
	ReadAll(bufSize int) (string, error)
	BytesRead() int64
}

type CountingToLowerReaderImpl struct {
	Reader         io.Reader
	TotalBytesRead int64
}

func NewCountingReader(r io.Reader) *CountingToLowerReaderImpl {
	return &CountingToLowerReaderImpl{
		Reader: r,
	}
}

func (cr *CountingToLowerReaderImpl) BytesRead() int64 {
	return cr.TotalBytesRead
}

func (cr *CountingToLowerReaderImpl) Read(p []byte) (int, error) {
	n, err := cr.Reader.Read(p)
	cr.TotalBytesRead += int64(n)

	// Convert uppercase English letters to lowercase
	for i := 0; i < n; i++ {
		if unicode.IsUpper(rune(p[i])) {
			// 65 (A) - 90 (Z), (a) 97 до (z) 122.
			// plus 32
			p[i] += 'a' - 'A'

		}
	}

	return n, err
}

func (cr *CountingToLowerReaderImpl) ReadAll(bufSize int) (string, error) {
	buf := make([]byte, bufSize)
	var content strings.Builder

	for {
		n, err := cr.Read(buf)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		content.Write(buf[:n])
	}

	return content.String(), nil
}
