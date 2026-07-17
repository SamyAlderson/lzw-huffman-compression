// Package main est le fichier principal du projet lzw-huffman-compression.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/your-username/golang-lzw"
	"github.com/your-username/golang-huffman"
)

// Args représente les arguments de la ligne de commande.
type Args struct {
	InputFile string
	OutputFile string
}

func main() {
	// Définition des arguments de la ligne de commande
	var args Args

	// Gestion des arguments
	if len(os.Args) != 3 {
		log.Fatal("Usage: main <input_file> <output_file>")
	}

	args.InputFile = os.Args[1]
	args.OutputFile = os.Args[2]

	// Lecture du fichier d'entrée
	data, err := os.ReadFile(args.InputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Compression du fichier
	compData, err := compress(data)
	if err != nil {
		log.Fatal(err)
	}

	// Écriture du fichier de sortie
	if err := os.WriteFile(args.OutputFile, compData, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fichier compressé en %d octets\n", len(compData))
}

// compress effectue la compression du fichier en utilisant l'algorithme LZW et Huffman.
func compress(data []byte) ([]byte, error) {
	// Encodage du fichier en utilisant l'algorithme LZW
	lzwData, err := lzw.Encode(data)
	if err != nil {
		return nil, err
	}

	// Création d'un arbre de Huffman à partir du fichier encodé
	huffmanTree, err := huffman.BuildTree(lzwData)
	if err != nil {
		return nil, err
	}

	// Encodage de l'arbre de Huffman en utilisant l'algorithme Huffman
	huffmanEnc, err := huffman.Encode(huffmanTree, lzwData)
	if err != nil {
		return nil, err
	}

	return huffmanEnc, nil
}