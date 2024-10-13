package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/hex"
)


// calculateSHA256 calculates the SHA256 hash of the given data
func calculateSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}


func main() {
	// First transaction
	s1 := "Hello world!"
	fmt.Printf("Message: %s\n", s1)
	fmt.Printf("Message: %s\n", time.Now().String())
	fmt.Printf("Message: %s\n", calculateSHA256(s1))
	fmt.Printf("Message: %s\n", calculateSHA256(s1+"!"))

	for i, n := range []int{10, 20, 30} {
		fmt.Printf("%d, %d\n", i, n)
	}
	
}
