.PHONY: bindings build test clean

CONTRACTS_OUT := ../muri-contracts/out
ABI_DIR := .abi
BINDINGS_DIR := chain/bindings

# Extract ABIs from Foundry output and generate Go bindings
bindings:
	@mkdir -p $(ABI_DIR)
	jq '.abi' $(CONTRACTS_OUT)/Market.sol/FileMarket.json > $(ABI_DIR)/FileMarket.json
	jq '.abi' $(CONTRACTS_OUT)/NodeStaking.sol/NodeStaking.json > $(ABI_DIR)/NodeStaking.json
	abigen --abi $(ABI_DIR)/FileMarket.json --pkg bindings --type FileMarket --out $(BINDINGS_DIR)/filemarket.go
	abigen --abi $(ABI_DIR)/NodeStaking.json --pkg bindings --type NodeStaking --out $(BINDINGS_DIR)/nodestaking.go
	@echo "Bindings generated."

build:
	go build ./...

test:
	go test ./... -v

clean:
	rm -rf $(ABI_DIR)
	go clean ./...
