# ByteSplitter

## Introduction

The `bytesplitter` package is a Go library designed to split slices of bytes into smaller slices based on a given memory size. It supports various units of memory like Kilobytes (KB) and Megabytes (MB).

## Features

- Types for memory sizes: `KB`, `MB`.
- Function for splitting byte slices: `Split`.

## Installation

```bash
go get -u github.com/alesr/bytesplitter
```

## Usage

### Importing the Package

```go
import "github.com/alesr/bytesplitter"
```

### Example 1: Splitting a Byte Slice into 1KB Chunks

```go
package main

import (
    "github.com/alesr/bytesplitter"
)

func main() {
    data := make([]byte, 5000) // 5KB of data

    // Split the data into 1KB chunks
    chunks, err := bytesplitter.Split(data, bytesplitter.KB(1))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Number of chunks:", len(chunks))
}
```

### Example 2: Splitting a Byte Slice into 2MB Chunks

```go
package main

import (
    "github.com/alesr/bytesplitter"
)

func main() {
    data := make([]byte, 5000000) // ~5MB of data

    // Split the data into 2MB chunks
    chunks, err := bytesplitter.Split(data, bytesplitter.MB(2))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Number of chunks:", len(chunks))
}
```

## API Documentation

### Types

- `KB`: Represents Kilobytes.
- `MB`: Represents Megabytes.

### Functions

- `Split(data []byte, size T) ([][]byte, error)`: Splits a byte slice `data` into smaller slices based on the given `size`. Returns an error if the splitting fails.

## Contributing

Feel free to submit issues and pull requests.
