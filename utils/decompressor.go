// Package decompressor contient les fonctions de décompression utilisant l'algorithme LZW et Huffman.
package decompressor

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/your-username/golang-lzw"
	"github.com/your-username/golang-huffman"
)

// Decompressor est une structure qui représente le décompresseur LZW et Huffman.
type Decompressor struct {
	lzwDecoder *lzw.Decoder
	huffmanDecoder *huffman.Decoder
}

// NewDecompressor retourne un nouveau décompresseur LZW et Huffman.
func NewDecompressor() *Decompressor {
	return &Decompressor{
		lzwDecoder: lzw.NewDecoder(),
		huffmanDecoder: huffman.NewDecoder(),
	}
}

// DecompressLZW décomprésser des données en utilisant l'algorithme LZW.
func (d *Decompressor) DecompressLZW(data []byte) ([]byte, error) {
	return d.lzwDecoder.Decode(data)
}

// DecompressHuffman décomprésser des données en utilisant l'algorithme Huffman.
func (d *Decompressor) DecompressHuffman(data []byte) ([]byte, error) {
	return d.huffmanDecoder.Decode(data)
}

// Decompress décomprésser des données en utilisant l'algorithme LZW et Huffman.
func (d *Decompressor) Decompress(data []byte) ([]byte, error) {
	// Décompresser les données en utilisant l'algorithme LZW
	lzwData, err := d.DecompressLZW(data)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la décompression LZW : %w", err)
	}

	// Décompresser les données en utilisant l'algorithme Huffman
	huffmanData, err := d.DecompressHuffman(lzwData)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la décompression Huffman : %w", err)
	}

	return huffmanData, nil
}

// DecompressWithLength décomprésser des données en utilisant l'algorithme LZW et Huffman, avec la longueur en huit bits.
func (d *Decompressor) DecompressWithLength(data []byte) ([]byte, error) {
	// Récupérer la longueur en huit bits
	length := binary.BigEndian.Uint32(data[:4])

	// Décompresser les données en utilisant l'algorithme LZW
	lzwData, err := d.DecompressLZW(data[4 : 4+length])
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la décompression LZW : %w", err)
	}

	// Décompresser les données en utilisant l'algorithme Huffman
	huffmanData, err := d.DecompressHuffman(lzwData)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la décompression Huffman : %w", err)
	}

	return huffmanData, nil
}