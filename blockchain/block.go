package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	BlockID      int
	PreviousHash string
	Transactions []Transaction
	Timestamp    time.Time
	Hash         string
}

// CreateBlock creates a new block with the given parameters
func CreateBlock(blockID int, previousHash string, transactions []Transaction) *Block {
	// TODO: Implement the CreateBlock function
	// 1. Create a new Block instance with the given parameters
	// 2. Set the Timestamp to the current time
	// 3. Calculate the Hash of the block using the CalculateHash method
	// 4. Return a pointer to the new Block
	return nil
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() string {
	// TODO: Implement the CalculateHash function
	// 1. Create a string that combines the block's data (BlockID, PreviousHash, Transactions, and Timestamp)
	// 2. Use the calculateSHA256 function to generate the hash of this string
	// 3. Return the calculated hash
	return ""
}

// calculateSHA256 calculates the SHA256 hash of the given data
func calculateSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

