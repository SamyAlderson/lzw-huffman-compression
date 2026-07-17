// Package lzw implémente l'algorithme LZW pour la compression de données.
package lzw

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// LZWCompressor est un objet qui implémente l'algorithme LZW pour la compression de données.
type LZWCompressor struct {
	buf      []byte
	codeSize int
}

// NewLZWCompressor crée un nouvel objet LZWCompressor.
func NewLZWCompressor(codeSize int) (*LZWCompressor, error) {
	if codeSize <= 0 {
		return nil, errors.New("codeSize doit être supérieur à 0")
	}
	return &LZWCompressor{codeSize: codeSize}, nil
}

// Compress compresse les données à l'aide de l'algorithme LZW.
func (c *LZWCompressor) Compress(r io.Reader) ([]byte, error) {
	c.buf = make([]byte, 0, 1024)

	var code int
	var prev byte

	for {
		n, err := r.Read(&c.buf)
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("erreur de lecture : %w", err)
			}
			break
		}

		for i := 0; i < n; i++ {
			curr := c.buf[i]
			if curr == prev {
				c.addCode(code)
				code = 256
			} else {
				code = c.incrementCode(code)
				c.addChar(curr)
			}
			prev = curr
		}
	}

	return c.buf, nil
}

func (c *LZWCompressor) addCode(code int) {
	c.buf = append(c.buf, byte(code))
}

func (c *LZWCompressor) incrementCode(code int) int {
	return code + 1
}

func (c *LZWCompressor) addChar(ch byte) {
	if c.codeSize > 12 {
		fmt.Println("CodeSize trop grand")
		os.Exit(1)
	}
	c.buf = append(c.buf, ch)
}

// Decompress décompresse les données à l'aide de l'algorithme LZW.
func Decompress(data []byte) ([]byte, error) {
	c := NewLZWCompressor(10)
	if c == nil {
		return nil, errors.New("erreur de création du LZWCompressor")
	}

	var code int
	var output []byte

	for i := 0; i < len(data); i++ {
		if data[i]&0x80 == 0x80 {
			code = int(data[i]) | ((int(data[i+1]) << 8) & 0xFF00)
			i++
		} else {
			code = int(data[i])
		}

		output = append(output, byte(code))
	}

	return output, nil
}