package blockchain

import "fmt"

// Transaction represents a transaction in the blockchain
type Transaction struct {
	Type      string
	AccountID string
	Amount    float64
}

// CreateTransaction creates a new Transaction
func CreateTransaction(transactionType, accountID string, amount float64) Transaction {
	// TODO: Implement the CreateTransaction function
	// 1. Create and return a new Transaction instance with the given parameters
	newTransaction := Transaction{
		Type: transactionType,
		AccountID: accountID,
		Amount: amount,
	}
	return newTransaction
}

// ValidateTransaction checks if the transaction is valid based on its type and amount
func ValidateTransaction(t Transaction) bool {
	// TODO: Implement the ValidateTransaction function
	// 1. Check the transaction type (deposit, withdraw, or transfer)
	if t.Type == "deposit" || t.Type == "withdraw" || t.Type == "transfer" {
		// Do nothing
	} else {
		return false
	}
	// 2. Ensure the amount is positive for all transaction types
	if t.Amount <= 0 {
		return false
	} else {
		// Do nothing
	}
	// 3. Return true if the transaction is valid, false otherwise
	return true
}

// String returns a string representation of the Transaction
func (t Transaction) String() string {
	return fmt.Sprintf("{Type: %s, AccountID: %s, Amount: %.2f}", t.Type, t.AccountID, t.Amount)
}

