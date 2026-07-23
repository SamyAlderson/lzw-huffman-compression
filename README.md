# LZW-Huffman Compression
LZW and Huffman in a single Go package

## What it does

This project implements both LZW and Huffman data compression algorithms. The main goal is to provide a simple and efficient way to compress data using these two popular compression techniques.

## Install

To install, run the following command in your terminal:
```bash
go get github.com/SamyAlderson/lzw-huffman-compression
```

## Usage

To use the package, import it in your Go file:
```go
import "github.com/SamyAlderson/lzw-huffman-compression"
```
Then, you can use the functions to compress and decompress data:
```go
import (
    "fmt"
    "github.com/SamyAlderson/lzw-huffman-compression"
)

func main() {
    // Compress data
    compress := lzwhuff Compression {
        LZW: true,
        Huffman: true,
    }
    data := []byte("Hello, World!")
    compressed, err := compress.Compress(data)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(compressed))

    // Decompress data
    decompress := lzwhuff Decompression {
        LZW: true,
        Huffman: true,
    }
    decompressed, err := decompress.Decompress(compressed)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(decompressed))
}
```

## Build from Source

To build from source, clone the repository:
```bash
git clone https://github.com/SamyAlderson/lzw-huffman-compression.git
```
Then, navigate to the project directory and run:
```bash
go build
```
This will build the package in the current directory.

## Project Structure

* `main.go`: the main package file
* `lzw.go`: LZW compression and decompression implementation
* `huffman.go`: Huffman compression and decompression implementation
* `compression.go`: the `Compression` struct and related functions
* `decompression.go`: the `Decompression` struct and related functions

## License

Copyright (c) 2026 SamyAlderson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.