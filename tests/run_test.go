package tests

import (
	"testing"
	"fmt"
	"strconv"
	bc "blockchain-api/blockchain"
)

func TestCreateGenesisBlock(t *testing.T) {
	tx1 := bc.Transaction{Type: "deposit", AccountID: "4321", Amount: 5000}
	genesisBlock := bc.CreateGenesisBlock([]bc.Transaction{tx1})
	if genesisBlock.BlockID != 1 {
		t.Errorf("Expected genesis block id to be 1, got %d", genesisBlock.BlockID)
	}
	if genesisBlock.PreviousHash != "0" {
		t.Errorf("Expected genesis block PrevHash to be '0', got %s", genesisBlock.PreviousHash)
	}
}

func TestAddNewBlock(t *testing.T) {
	// First transaction
	tx1 := bc.Transaction{Type: "deposit", AccountID: "1234", Amount: 2000}
	
	// Create empty Blockchain
	var blockchain bc.Blockchain
	
	// Add transaction to blockchain
	blockchain.AddNewBlock([]bc.Transaction{tx1})

	// Another transaction
	tx2 := bc.Transaction{Type: "withdraw", AccountID: "1234", Amount: 500}

	// Add to blockchain
	blockchain.AddNewBlock([]bc.Transaction{tx2})

	genesisBlock := blockchain.Blocks[0]
	nextBlock := blockchain.Blocks[1]

	if genesisBlock.BlockID != 1 || genesisBlock.PreviousHash != "0"  {
		t.Errorf("Expected block id to be 1, got %d", genesisBlock.BlockID)
	}
	if genesisBlock.Hash != nextBlock.PreviousHash {
		t.Errorf("Expected new block's PrevHash to be %s, got %s", genesisBlock.Hash, nextBlock.PreviousHash)
	}
}

func TestHashing(t *testing.T) {
	tx1 := bc.Transaction{Type: "deposit", AccountID: "4321", Amount: 5000}
	genesisBlock := bc.CreateGenesisBlock([]bc.Transaction{tx1})
	blockData := strconv.Itoa(genesisBlock.BlockID) + genesisBlock.PreviousHash + genesisBlock.Timestamp.String() + fmt.Sprintf("%v", genesisBlock.Transactions)
	calculatedHash := bc.CalculateSHA256(blockData)
	if genesisBlock.Hash != calculatedHash {
		t.Errorf("Expected block hash to be %s, got %s", calculatedHash, genesisBlock.Hash)
	}
}

func TestBlockchainImmutability(t *testing.T) {
	// First transaction
	tx1 := bc.Transaction{Type: "deposit", AccountID: "1234", Amount: 2000}
	
	// Create empty Blockchain
	var blockchain bc.Blockchain
	
	// Add transaction to blockchain
	blockchain.AddNewBlock([]bc.Transaction{tx1})

	// Another transaction
	tx2 := bc.Transaction{Type: "withdraw", AccountID: "1234", Amount: 500}

	// Add to blockchain
	blockchain.AddNewBlock([]bc.Transaction{tx2})

	// Test for immutability (tampering with data should fail validation)
	blockchain.Blocks[1].Transactions = []bc.Transaction{bc.Transaction{Type: "deposit", AccountID: "4321", Amount: 1000}}
	if bc.ValidateBlockchainIntegrity(blockchain) {
		t.Errorf("Blockchain validation failed. Tampering with data should invalidate the chain.")
	}
}

func TestBlockchainValidation(t *testing.T) {
	// First transaction
	tx1 := bc.Transaction{Type: "deposit", AccountID: "1234", Amount: 2000}
	
	// Create empty Blockchain
	var blockchain bc.Blockchain
	
	// Add transaction to blockchain
	blockchain.AddNewBlock([]bc.Transaction{tx1})

	// Another transaction
	tx2 := bc.Transaction{Type: "withdraw", AccountID: "1234", Amount: 500}

	// Add to blockchain
	blockchain.AddNewBlock([]bc.Transaction{tx2})

	// Validate Blockchain Integrity
	if bc.ValidateBlockchainIntegrity(blockchain) != true {
		t.Errorf("Expected blockchain to be valid")
	}

	// Simulate tampering with a block
	blockchain.Blocks[1].Hash = "123456"
	if bc.ValidateBlockchainIntegrity(blockchain) {
		t.Errorf("Expected blockchain to be invalid after tampering with a block")
	}
}
