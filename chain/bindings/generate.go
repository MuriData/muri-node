package bindings

//go:generate sh -c "jq '.abi' ../../.abi/FileMarket.json > /tmp/filemarket_abi.json && abigen --abi /tmp/filemarket_abi.json --pkg bindings --type FileMarket --out filemarket.go"
//go:generate sh -c "jq '.abi' ../../.abi/NodeStaking.json > /tmp/nodestaking_abi.json && abigen --abi /tmp/nodestaking_abi.json --pkg bindings --type NodeStaking --out nodestaking.go"
