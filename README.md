# lzw-huffman-compression
[![Go](https://img.shields.io/badge/Langue-Go-blue)](https://golang.org/)
[![Licence](https://img.shields.io/badge/Licence-MIT-yellowgreen)](https://opensource.org/licenses/MIT)
[![CI](https://github.com/username/lzw-huffman-compression/actions/workflows/ci.yml/badge.svg)](https://github.com/username/lzw-huffman-compression/actions/workflows/ci.yml)

## Description

Ce projet implémente les algorithmes LZW et Huffman pour la compression de données. LZW (Lempel-Ziv-Welch) est un algorithme de compression de données qui utilise une table de codes pour représenter les séquences de caractères répétitives. Huffman est un algorithme de codage arithmétique qui utilise une arbre binaire pour représenter les symboles les plus fréquents.

## Fonctionnalités

* Compression de données en utilisant l'algorithme LZW
* Compression de données en utilisant l'algorithme Huffman
* Décompression de données compressées
* Utilisation de la bibliothèque Go pour une implémentation efficace

## Installation

Pour installer ce projet, vous devez avoir Go installé sur votre système. Vous pouvez suivre les étapes suivantes pour installer Go :
```bash
# Sous Linux/Mac
curl -O https://dl.google.com/go/go1.19.4.linux-amd64.tar.gz
tar -xvf go1.19.4.linux-amd64.tar.gz
sudo mv go /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Sous Windows
# Télécharger l'installateur de Go de la page officielle
# Suivre les instructions d'installation
```
Une fois Go installé, vous pouvez cloner ce projet à l'aide de Git :
```bash
git clone https://github.com/username/lzw-huffman-compression.git
cd lzw-huffman-compression
```
## Usage avec exemples

Pour utiliser ce projet, vous devez exécuter les commandes suivantes :
```bash
go build cmd/lzw-huffman/main.go
./lzw-huffman -c "Bonjour, monde !" -o output.lzw
./lzw-huffman -d output.lzw
```
Cela compressera le texte "Bonjour, monde !" dans le fichier `output.lzw` et décomposera le fichier `output.lzw` pour afficher le texte original.

## Architecture du projet

Le projet est composé des fichiers suivants :

* `cmd/lzw-huffman/main.go` : Fichier principal qui gère les commandes d'utilisation
* `utils/lzw.go` : Implémentation de l'algorithme LZW
* `utils/huffman.go` : Implémentation de l'algorithme Huffman
* `utils/compressor.go` : Compresseur utilisant LZW et Huffman
* `utils/decompressor.go` : Décompresseur utilisant LZW et Huffman

## Contribuer

Si vous souhaitez contribuer à ce projet, vous pouvez suivre les étapes suivantes :

1. Cloner le projet à l'aide de Git
2. Créer une branche pour votre contribution
3. Effectuer les modifications nécessaires
4. Soumettre une demande de pull-request

## Licence

Ce projet est sous licence MIT. Vous pouvez télécharger la licence à l'adresse suivante : https://opensource.org/licenses/MIT.