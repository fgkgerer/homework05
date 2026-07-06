# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
# Download dependencies
go mod download

# Build
go build -o task2.exe .

# Run (requires ETH_RPC_URL env var pointing to a Sepolia testnet RPC endpoint)
ETH_RPC_URL=<rpc-url> go run .

# Run tests (if any)
go test ./...
```

## Architecture

This is a Go homework project for interacting with Ethereum smart contracts on the Sepolia testnet using `go-ethereum` (v1.17.4).

**Tooling:** `abigen` (from go-ethereum) is used to generate Go binding code from a Solidity contract's ABI and bytecode. The generated bindings are then used to call contract methods from Go.

**Workflow:**
1. Write a Solidity contract (e.g., a counter) and compile it to ABI + bytecode
2. Run `abigen` to generate Go bindings from those artifacts
3. Use the generated code + `ethclient` to connect and interact with the deployed contract on Sepolia

**Key dependency:** `github.com/ethereum/go-ethereum/ethclient` — provides the Ethereum JSON-RPC client for connecting to nodes.

**Configuration:** The RPC endpoint is passed via the `ETH_RPC_URL` environment variable. No private key management is implemented yet.
