package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"strconv"
)

// Block represents a block in the blockchain
type Block struct {
	BlockID      int
	PreviousHash string
	Transactions []Transaction
	Timestamp    time.Time
	Hash         string
}

// Blockchain represents the blockchain
type Blockchain struct {
	Blocks []Block
}

// CreateBlock creates a new block with the given parameters
func CreateBlock(blockID int, previousHash string, transactions []Transaction) Block {
	// TODO: Implement the CreateBlock function
	// 1. Create a new Block instance with the given parameters
	// 2. Set the Timestamp to the current time
	// 3. Calculate the Hash of the block using the CalculateHash method
	// 4. Return the new Block
}

// GenesisBlock creates a genesis block with the given parameters
func CreateGenesisBlock(transactions []Transaction) Block {
	// TODO: Implement the CreateGenesisBlock function
	// 1. Create a new Block instance with the given transactions, BlockID=1, PreviousHash="0"
	// 2. Set the Timestamp to the current time
	// 3. Calculate the Hash of the block using the CalculateHash method
	// 4. Return the new Block
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() string {
	// TODO: Implement the CalculateHash function
	// 1. Create a string that combines the block's data (BlockID, PreviousHash, Transactions, and Timestamp)
	//    - To convert integer to string use strconv.Itoa(...)
	// 2. Use the calculateSHA256 function to generate the hash of this string
	// 3. Set the block's Hash to calculated hash
}

// CalculateSHA256 calculates the SHA256 hash of the given data
func CalculateSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}


func (bchain *Blockchain) AddNewBlock(transactions []Transaction) {
	// TODO: Implement the AddNewBlock function
	// If the blockchain is empty, then add a Genesis block to it with the given transactions
	// If the blockchain is not empty, then add a new block to it with the given transactions
}


func ValidateBlockchainIntegrity(bchain Blockchain) bool {
	// If blockchain is empty, return true

	// Validate Hash and PreviousHash of all blocks in blockchain:
	//   The PreviousHash of the current block in blockchain should match the Hash of previous block in blockchain. Else return false.

	// Validate all transactions within a block in blockchain:
	//   Use ValidateTransaction function to validate transactions of each block in blockchain. If any transaction is invalid return false.

	return true
}

