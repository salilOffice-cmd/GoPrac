package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset represents an asset in the ledger.
type Asset struct {
	ID     string `json:"id"`
	Owner  string `json:"owner"`
	Color  string `json:"color"`
	Size   int    `json:"size"`
	Price  int    `json:"price"`
	Status string `json:"status"`
}

// SimpleAssetChaincode defines the Smart Contract structure.
type SimpleAssetChaincode struct {
	contractapi.Contract
}

// CreateAsset creates a new asset in the ledger.
func (s *SimpleAssetChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, assetJSON string) error {
	// Unmarshal the asset JSON into an Asset struct
	var asset Asset
	err := json.Unmarshal([]byte(assetJSON), &asset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal asset JSON: %v", err)
	}

	// Validate asset fields
	if asset.ID == "" {
		return fmt.Errorf("asset ID is required")
	}
	if asset.Owner == "" {
		return fmt.Errorf("asset owner is required")
	}
	if asset.Price <= 0 {
		return fmt.Errorf("asset price must be positive")
	}

	// Check if asset already exists
	exists, err := s.AssetExists(ctx, asset.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("asset with ID %s already exists", asset.ID)
	}

	// Store asset in the ledger
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset JSON: %v", err)
	}
	err = ctx.GetStub().PutState(asset.ID, assetBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset state: %v", err)
	}

	return nil
}

// AssetExists checks if an asset with the given ID exists in the ledger.
func (s *SimpleAssetChaincode) AssetExists(ctx contractapi.TransactionContextInterface, assetID string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(assetID)
	if err != nil {
		return false, fmt.Errorf("failed to read asset state: %v", err)
	}
	return assetBytes != nil, nil
}

// Main function
func main() {
	chaincode, err := contractapi.NewChaincode(&SimpleAssetChaincode{})
	if err != nil {
		fmt.Printf("Error creating SimpleAssetChaincode: %v\n", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting SimpleAssetChaincode: %v\n", err)
	}
}
