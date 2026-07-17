// Package huffman contient l'implémentation de l'algorithme Huffman.
package huffman

import (
	"container/heap"
	"fmt"
	"math"
)

// Node représente un nœud dans l'arbre de Huffman.
type Node struct {
	Caractere byte
	Frequency int
	Left     *Node
	Right    *Node
}

// priorityQueue est une pile utilisée pour trier les nœuds en fonction de leur fréquence.
type priorityQueue []*Node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Frequency < pq[j].Frequency
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Node))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// BuildHuffmanTree construit l'arbre de Huffman à partir d'une liste de fréquences.
func BuildHuffmanTree(frequency map[byte]int) (*Node, error) {
	// Créer une liste de nœuds pour chaque caractère
	pq := make(priorityQueue, 0)
	for char, freq := range frequency {
		node := &Node{
			Caractere: char,
			Frequency: freq,
		}
		heap.Push(&pq, node)
	}

	// Tant qu'il reste plus d'un nœud dans la pile
	for pq.Len() > 1 {
		// Extraire les deux nœuds les moins fréquents
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		// Créer un nouveau nœud qui est la fusion des deux précédents
		parent := &Node{
			Frequency: left.Frequency + right.Frequency,
			Left:     left,
			Right:    right,
		}
		heap.Push(&pq, parent)
	}

	// Le nœud racine est le seul nœud restant dans la pile
	root := pq[0]

	return root, nil
}

// Encode utilise l'arbre de Huffman pour encoder les données.
func Encode(root *Node, data []byte) ([]byte, error) {
	var codeMap = make(map[byte]string)
	var encode func(node *Node, code string)
	encode = func(node *Node, code string) {
		if node != nil {
			encode(node.Left, code + "0")
			encode(node.Right, code + "1")
			codeMap[node.Caractere] = code
		}
	}
	encode(root, "")

	var encodedData []byte
	for _, char := range data {
		encodedData = append(encodedData, []byte(codeMap[char])...)
	}

	return encodedData, nil
}

// Decode utilise l'arbre de Huffman pour décoder les données.
func Decode(root *Node, encodedData []byte) ([]byte, error) {
	var decode func(node *Node, code string) byte
	decode = func(node *Node, code string) byte {
		if node == nil {
			return 0
		}
		if code == "" {
			return node.Caractere
		}
		if code[0] == '0' {
			return decode(node.Left, code[1:])
		}
		return decode(node.Right, code[1:])
	}

	var decodedData []byte
	for _, code := range encodedData {
		decodedData = append(decodedData, byte(decode(root, string(code))))
	}

	return decodedData, nil
}

// calculateHuffmanCodes utilise l'arbre de Huffman pour calculer les codes de Huffman.
func calculateHuffmanCodes(root *Node) (map[byte]string, error) {
	var codeMap = make(map[byte]string)
	var traverse func(node *Node, code string)
	traverse = func(node *Node, code string) {
		if node != nil {
			traverse(node.Left, code + "0")
			traverse(node.Right, code + "1")
			codeMap[node.Caractere] = code
		}
	}
	traverse(root, "")

	return codeMap, nil
}

// calculateHuffmanTree utilise l'arbre de Huffman pour calculer la fréquence de chaque nœud.
func calculateHuffmanTree(root *Node) (map[*Node]int, error) {
	var frequency = make(map[*Node]int)
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node != nil {
			frequency[node]++
			traverse(node.Left)
			traverse(node.Right)
		}
	}
	traverse(root)

	return frequency, nil
}

func main() {
	// Exemple d'utilisation
	frequency := map[byte]int{'A': 10, 'B': 5, 'C': 5}
	root, err := BuildHuffmanTree(frequency)
	if err != nil {
		fmt.Println(err)
		return
	}

	encodedData, err := Encode(root, []byte("ABC"))
	if err != nil {
		fmt.Println(err)
		return
	}

	decodedData, err := Decode(root, encodedData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Encoded data:", encodedData)
	fmt.Println("Decoded data:", decodedData)
}
```

Ce code définit les fonctions suivantes :

- `BuildHuffmanTree(frequency map[byte]int) (*Node, error)`: construit l'arbre de Huffman à partir d'une liste de fréquences.
- `Encode(root *Node, data []byte) ([]byte, error)`: utilise l'arbre de Huffman pour encoder les données.
- `Decode(root *Node, encodedData []byte) ([]byte, error)`: utilise l'arbre de Huffman pour décoder les données.
- `calculateHuffmanCodes(root *Node) (map[byte]string, error)`: utilise l'arbre de Huffman pour calculer les codes de Huffman.
- `calculateHuffmanTree(root *Node) (map[*Node]int, error)`: utilise l'arbre de Huffman pour calculer la fréquence de chaque nœud.
