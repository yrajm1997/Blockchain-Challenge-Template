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
	newBlock := Block{
		BlockID: blockID,
		PreviousHash: previousHash,
		Transactions: transactions,
	}
	// 2. Set the Timestamp to the current time
	newBlock.Timestamp = time.Now()
	// 3. Calculate the Hash of the block using the CalculateHash method
	newBlock.CalculateHash()
	// 4. Return a pointer to the new Block
	return newBlock
}

// GenesisBlock creates a genesis block with the given parameters
func CreateGenesisBlock(transactions []Transaction) Block {
	// TODO: Implement the CreateBlock function
	// 1. Create a new Block instance with the given parameters
	newBlock := Block{
		BlockID: 1,
		PreviousHash: "0",
		Transactions: transactions,
	}
	// 2. Set the Timestamp to the current time
	newBlock.Timestamp = time.Now()
	// 3. Calculate the Hash of the block using the CalculateHash method
	newBlock.CalculateHash()
	// 4. Return a pointer to the new Block
	return newBlock
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() string {
	// TODO: Implement the CalculateHash function
	// 1. Create a string that combines the block's data (BlockID, PreviousHash, Transactions, and Timestamp)
	record := strconv.Itoa(b.BlockID) + b.PreviousHash + b.Timestamp.String() +  fmt.Sprintf("%v", b.Transactions)
	// 2. Use the calculateSHA256 function to generate the hash of this string
	hashed := CalculateSHA256(record)
	// 3. Return the calculated hash
	b.Hash = hashed
	return hashed
}

// calculateSHA256 calculates the SHA256 hash of the given data
func CalculateSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}


func (bchain *Blockchain) AddNewBlock(transactions []Transaction) {
	if len(bchain.Blocks) == 0 {
		newBlock := CreateGenesisBlock(transactions)
		bchain.Blocks = append(bchain.Blocks, newBlock)
	} else {
		prev_block := bchain.Blocks[len(bchain.Blocks)-1]
		newBlock := CreateBlock(prev_block.BlockID + 1, prev_block.Hash, transactions)
		bchain.Blocks = append(bchain.Blocks, newBlock)
	}
}


func ValidateBlockchainIntegrity(bchain Blockchain) bool {
	// If blockchain is empty
	if len(bchain.Blocks) == 0 {
		fmt.Printf("Blockchain is empty!\n")
		return true
	}
	// Validate Hash and PreviousHash of all blocks in blockchain
	for i, block := range bchain.Blocks {
		if i > 0 {
			if bchain.Blocks[i-1].Hash != block.PreviousHash {
				fmt.Printf("Block being added is invalid!\n")
				return false
			}
		} else {
			if block.PreviousHash != "0" {
				fmt.Printf("Genesis Block's previous hash is wrong!\n")
				return false
			}
		}
		blockData := strconv.Itoa(block.BlockID) + block.PreviousHash + block.Timestamp.String() + fmt.Sprintf("%v", block.Transactions)
		if block.Hash != CalculateSHA256(blockData) {
			fmt.Printf("In block %d, the hash value is invalid!\n", i)
			return false
		}
	}
	// Validate all transactions within a block in blockchain
	for i, block := range bchain.Blocks {
		for j, transaction := range block.Transactions {
			if ValidateTransaction(transaction) == false {
				fmt.Printf("In block %d, at position %d, the transaction is invalid!\n%s\n", i, j, transaction)
				return false
			}
		}
	}
	return true
}

