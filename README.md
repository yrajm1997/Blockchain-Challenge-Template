# Blockchain Development Challenge

An implementation of core blockchain concepts using Go, with GitHub Codespaces integration.

## Objective
Your task is to complete the implementation of a basic blockchain system. You'll work on creating blocks, handling transactions, and maintaining the integrity of the blockchain.

## Getting Started
1. **Fork this project** to create your own copy of the repository.
2. **Use GitHub Codespaces**:
   - Click on the green **Code** button in your forked repository.
   - Select **Codespaces** and choose "Create codespace on main" to open your development environment.

3. The repository includes a pre-configured `go.yml` file, which automatically sets up the Go environment in Codespaces. 
   - **Ensure the Go version is correct**: Check that the `go.yml` file has the correct Go version for the project.
     You can update the `go.yml` file if necessary:
     ```yaml
     - name: Set up Go
       uses: actions/setup-go@v2
       with:
         go-version: '1.16'  # Adjust this version based on your project's needs
     ```

4. Once the Codespace is ready and the environment is set up, review the code in the `blockchain/` folder and `main.go` to understand the structure.
5. Implement the missing functions marked with **TODO** comments.
6. Test your implementation by running the `main.go` file within GitHub Codespaces.

## Files to Work On
- `blockchain/block.go`: Implement block creation and hash calculation.
- `blockchain/transaction.go`: Implement transaction creation and validation.
- `main.go`: Create the blockchain, add transactions, and test your implementation.

## Requirements
- **`block.go`**:
  - Implement the `CreateBlock` function to define how a block is created.
  - Implement the `CalculateHash` function to calculate the hash of each block using the block's data.

- **`transaction.go`**:
  - Implement the `CreateTransaction` function to handle transaction creation.
  - Implement the `ValidateTransaction` function to check transaction validity.

- **`main.go`**:
  - Create the genesis block, add transactions, and create additional blocks in the blockchain.

## Steps to Follow
1. **Fork** the repository and set up your environment using **GitHub Codespaces**.
2. **block.go**: Implement block creation logic and hash calculation.
3. **transaction.go**: Implement transaction creation and validation.
4. **main.go**: Test the blockchain by adding a genesis block, creating transactions, and adding additional blocks.
5. Run the code inside **GitHub Codespaces** to test your implementation.

## Tips
- Use the provided `calculateSHA256` function in `block.go` for hash calculation.
- Ensure that each block's `PreviousHash` correctly points to the hash of the previous block to maintain chain integrity.
- Implement proper validation for transactions based on their type (deposit, withdraw, transfer) and amount.

## Submission Guidelines
After completing the challenge, submit the link to your **forked GitHub repository** in the **LMS submission link text box**.

## Evaluation Criteria
- Correct implementation of blockchain and block creation logic.
- Proper handling of transactions (deposit, withdraw, transfer).
- Accurate use of hashing to maintain blockchain integrity.
- Clean and readable code with appropriate comments and structure.
- Successful execution of the project within **GitHub Codespaces**.

Good luck, and happy coding!
