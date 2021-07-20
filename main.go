package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

/*
B1

	b1Hash = (data + "")

B2
	b2Hash = (data + b1Hash)

B3
	b3Hash = (data + b2Hash)
*/

func main() {
	genesisBlock := block{"Genesis Block", "", ""} // prevhash and has not exist
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))

	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}
