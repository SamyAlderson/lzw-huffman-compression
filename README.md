# LZW-Huffman Compression
Efficient compression using LZW and Huffman algorithms

This project implements the LZW and Huffman data compression algorithms in Go. It's a basic implementation for educational purposes, aiming to demonstrate the concepts behind these algorithms.

## Installation

```bash
go get github.com/SamyAlderson/lzw-huffman-compression
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/SamyAlderson/lzw-huffman-compression/compression"
)

func main() {
	data := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	compressed, err := compression.LZW(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(compressed)
}
```

## Building from Source

```bash
go build main.go
```

## Running Tests (none available)

## Project Structure

- `compression/lzw.go`: LZW compression algorithm implementation
- `compression/huffman.go`: Huffman compression algorithm implementation
- `main.go`: Example usage of the compression algorithms
- `README.md`: This file
- `LICENSE`: License information

## License

Copyright (c) 2026 SamyAlderson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.