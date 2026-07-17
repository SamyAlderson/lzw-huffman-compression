// Package compressor implémente un compresseur utilisant l'algorithme LZW et Huffman.
package compressor

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/your-username/golang-lzw"
	"github.com/your-username/golang-huffman"
)

// Compressor est un compresseur utilisant l'algorithme LZW et Huffman.
type Compressor struct {
	lzw   *lzw.LZW
	huffman *huffman.Huffman
}

// NewCompressor crée un nouveau compresseur.
func NewCompressor() *Compressor {
	return &Compressor{
		lzw:   lzw.NewLZW(),
		huffman: huffman.NewHuffman(),
	}
}

// Compress comprime les données à l'aide de l'algorithme LZW puis de Huffman.
func (c *Compressor) Compress(data []byte) ([]byte, error) {
	// Compression LZW
	lzwData, err := c.lzw.Compress(data)
	if err != nil {
		return nil, fmt.Errorf("compression LZW échouée: %w", err)
	}

	// Compression Huffman
	return c.huffman.Compress(lzwData)
}

// Decompress décompresse les données à l'aide de l'algorithme LZW puis de Huffman.
func (c *Compressor) Decompress(compressed []byte) ([]byte, error) {
	// Décompression Huffman
	huffmanData, err := c.huffman.Decompress(compressed)
	if err != nil {
		return nil, fmt.Errorf("décompression Huffman échouée: %w", err)
	}

	// Décompression LZW
	return c.lzw.Decompress(huffmanData)
}

// WriteBytes écrit les données compressées dans le buffer.
func (c *Compressor) WriteBytes(w io.Writer, compressed []byte) error {
	// Écriture des données compressées dans le buffer
	err := binary.Write(w, binary.LittleEndian, uint32(len(compressed)))
	if err != nil {
		return fmt.Errorf("écriture de la taille des données compressées échouée: %w", err)
	}

	// Écriture des données compressées elles-mêmes dans le buffer
	_, err = w.Write(compressed)
	if err != nil {
		return fmt.Errorf("écriture des données compressées échouée: %w", err)
	}

	return nil
}

// ReadBytes lit les données compressées du buffer.
func (c *Compressor) ReadBytes(r io.Reader) ([]byte, error) {
	// Lecture de la taille des données compressées
	var size uint32
	err := binary.Read(r, binary.LittleEndian, &size)
	if err != nil {
		return nil, fmt.Errorf("lecture de la taille des données compressées échouée: %w", err)
	}

	// Allocation d'un buffer pour les données compressées
	compressed := make([]byte, size)

	// Lecture des données compressées elles-mêmes
	_, err = io.ReadFull(r, compressed)
	if err != nil {
		return nil, fmt.Errorf("lecture des données compressées échouée: %w", err)
	}

	return compressed, nil
}