# Blockchain Development Playground Challenge: Banking Ledger System

An implementation of core blockchain concepts using Go.

## Objective
You are tasked with developing a simple blockchain-based banking ledger system using Go. The blockchain will record transactions such as deposits, withdrawals, and transfers between accounts, ensuring immutability and transparency of transactions. The system will include the creation of blocks, hashing, and validation mechanisms to ensure the security and consistency of the ledger.

## Steps to Follow
1. This GitHub repository will be accessible to you once you accept the Assignment.
2. You have to work directly in this GitHub repository. It is like your own copy of the original repository.

3. The folder structure is as given below:
   ```
   root
   ├── blockchain/
   │   ├── block.go
   │   └── transaction.go
   ├── tests/
   │   └── run_test.go
   ├── main.go
   └── go.mod
   ```

4. To edit and run the files **Use GitHub Codespaces**:
   - Click on the green **Code** dropdown button in this repository.
   - Select **Codespaces** tab and choose *"Create codespace on main"* to open your development environment.
   - After a minute or two, your Codespace will be ready.

5. In the Codespace IDE, *Go* comes pre-installed. 
   - **Check the Go version installed**: Run the following command in terminal `go version` to check the version.
   - **Ensure the same Go version is in `go.mod` file**: Check the `go.mod` file. It should have the same Go version which is pre-installed. 
     You can update the `go.mod` file if necessary:
     ```yaml
     module blockchain-api
     
     go 1.23.1
     ```

6. Review the code in the `blockchain/` folder and `main.go` to understand the structure.
7. Implement the missing functions marked with **TODO** comments.
8. Test your implementation by running the `main.go` file within GitHub Codespace.
   - **To run the application**: Run the following command in terminal `go run main.go`
   - **To run the test cases**: Run the following command in terminal `go test ./tests -v`

## Files to Work On
- `blockchain/block.go`
- `blockchain/transaction.go`

## Requirements
- **`block.go`**:
  - Implement the `CreateBlock` function to define how a block is created.
  - Implement the `CreateGenesisBlock` function to define how a genesis block is created.
  - Implement the `CalculateHash` function to calculate the hash of each block using the block's data.
  - Implement the `AddNewBlock` function to add a new block to the Blockchain.
  - Implement the `ValidateBlockchainIntegrity` function to check the Blockchain integrity.

- **`transaction.go`**:
  - Implement the `CreateTransaction` function to handle transaction creation.
  - Implement the `ValidateTransaction` function to check transaction validity.

- **`main.go`**:
  - Use this file to import and test the functions present in `block.go` and `transaction.go` files.


## Tips
- Use the provided `CalculateSHA256` function in `block.go` for hash calculation.
- Ensure that each block's `PreviousHash` correctly points to the 'Hash' of the previous block to maintain chain integrity.
- Implement proper validation for transactions based on their type (deposit, withdraw, transfer) and amount.

## Submission Guidelines
After completing the challenge and developing the solution code in GitHub Codespace, commit and push the changes to this repository. The latest push will be considered for grading.

## Evaluation Criteria
- Correct implementation of blockchain and block creation logic.
- Proper handling of transactions (deposit, withdraw, transfer).
- Accurate use of hashing to maintain blockchain integrity.
- Clean and readable code with appropriate comments and structure.
- Successful execution of the project within **GitHub Codespaces**.

Good luck, and happy coding!
