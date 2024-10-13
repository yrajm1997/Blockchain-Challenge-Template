package main

import (
	"fmt"
)


func main() {
	// First transaction
	tx1 := Transaction{Type: "deposit", AccountID: "1234", Amount: 2000}
	
	// Create empty Blockchain
	var blockchain Blockchain
	
	// Add transaction to blockchain
	blockchain.AddNewBlock([]Transaction{tx1})

	// Another transaction
	tx2 := Transaction{Type: "withdraw", AccountID: "1234", Amount: 500}

	// Add to blockchain
	blockchain.AddNewBlock([]Transaction{tx2})

	// Validate Blockchain Integrity
	fmt.Printf("Blockchain Integrity: %t\n", ValidateBlockchainIntegrity(blockchain))

	// Display the blockchain
	for _, block := range blockchain.Blocks {
		fmt.Printf("Block %d\n", block.BlockID)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Transactions: %v\n", block.Transactions)
		fmt.Printf("Timestamp: %s\n\n", block.Timestamp.String())
	}
}
