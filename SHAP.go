package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	input := []byte(os.Args[1])
	hash := sha256.Sum256(input)
	fmt.Printf("%x", hash)
}
