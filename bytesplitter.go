package bytesplitter

import "fmt"

// KB is a byte size in kilobytes.
type KB int

func (kb KB) toBytes() int {
	return int(kb) * 1024
}

// MB is a byte size in megabytes.
type MB int

func (mb MB) toBytes() int {
	return int(mb) * 1024 * 1024
}

// ConvertibleToByteSize is an interface for types that can be converted to byte sizes.
type ConvertibleToByteSize interface {
	toBytes() int
}

// Split splits a slice of bytes into smaller slices based on a given memory size.
func Split[T ConvertibleToByteSize](data []byte, size T) ([][]byte, error) {
	var chunks [][]byte

	if len(data) == 0 {
		return chunks, nil
	}

	chunkSize := int(size.toBytes())

	if chunkSize <= 0 {
		return nil, fmt.Errorf("invalid chunk size: %d", chunkSize)
	}

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		// Make sure not to go out of bounds
		if end > len(data) {
			end = len(data)
		}

		chunks = append(chunks, data[i:end])
	}
	return chunks, nil
}
