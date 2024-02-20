package main

// At the time of release and till hyperledger fabric 2.x,
// low level APIs are being used.
// But in the modern world, we now use contractapi which are also called as high level APIs
// In high level APIs, Init and Invoke functions are not essential
// However, we can define an Init function and we will invoke this function for instantiation of the 
// chaincode on a peer after the chaincode definition has been committed to the channel.
// In low level APIs, we had to use shim.ChaincodeStubInterface, detail code of low level api in chapter 2B

// To read more about contractApi, refer
// https://pkg.go.dev/github.com/hyperledger/fabric-contract-api-go/contractapi

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)



// You can say that SimpleAssetChaincode acts as the smart contract
// In other words, SimpleAssetChaincode defines the Smart Contract structure
// You can add methods to the struct by defining your custom functions
// So, in that way all your custom functions of the contract can be called using this struct
// The syntax below indicates that SimpleAssetChaincode inherits all the functions and properties 
// from contractapi.Contract p
type SimpleAssetChaincode struct {
	contractapi.Contract
}



// Asset represents a single asset
// This added will get stored in the ledger like this in the form of key value pairs:
// ID : Asset{"ID" : "1", "owner": "salil", "color" : "Red", ...}  (in json format)
type Asset struct {
	ID     string `json:"ID"`
	Owner  string `json:"owner"`
	Color  string `json:"color"`
	Size   int    `json:"size"`
	Price  int    `json:"price"`
}
// the `json:"color"` tells the chaincode that whenever a function receives a parameter of type Asset,
// unmarshal/deserialize it like shown above
// In short, this code maps the recevied 'json string' to a 'go struct'
// Also, when we do the vice versa meaning from 'go struct' to 'json string', the chaincode knows 
// how to create a json string from the go struct with the help of above code




// InitLedger adds a base set of assets to the ledger
func (s *SimpleAssetChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "asset1", Owner: "Alice", Color: "red", Size: 5, Price: 100},
		{ID: "asset2", Owner: "Bob", Color: "blue", Size: 10, Price: 200},
	}

	for _, asset := range assets {
		err := s.CreateAsset(ctx, asset)
		if err != nil {
			return fmt.Errorf("failed to create asset %s: %v", asset.ID, err)
		}
	}

	return nil
}

// CreateAsset adds a new asset to the ledger
func (s *SimpleAssetChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, asset Asset) error {
	exists, err := s.AssetExists(ctx, asset.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", asset.ID)
	}

	err = ctx.GetStub().PutState(asset.ID, []byte(asset.Owner))
	if err != nil {
		return fmt.Errorf("failed to create asset: %v", err)
	}

	return nil
}

// AssetExists checks if an asset exists in the ledger
func (s *SimpleAssetChaincode) AssetExists(ctx contractapi.TransactionContextInterface, assetID string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(assetID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetBytes != nil, nil
}

// QueryAsset returns the asset stored in the ledger
func (s *SimpleAssetChaincode) QueryAsset(ctx contractapi.TransactionContextInterface, assetID string) (*Asset, error) {
	assetBytes, err := ctx.GetStub().GetState(assetID)
	if err != nil {
		return nil, fmt.Errorf("failed to read asset %s from world state: %v", assetID, err)
	}
	if assetBytes == nil {
		return nil, fmt.Errorf("the asset %s does not exist", assetID)
	}

	asset := new(Asset)
	err = json.Unmarshal(assetBytes, asset)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return asset, nil
}

// UpdateAsset updates an existing asset in the ledger
func (s *SimpleAssetChaincode) UpdateAsset(ctx contractapi.TransactionContextInterface, asset Asset) error {
	exists, err := s.AssetExists(ctx, asset.ID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", asset.ID)
	}

	err = ctx.GetStub().PutState(asset.ID, []byte(asset.Owner))
	if err != nil {
		return fmt.Errorf("failed to update asset: %v", err)
	}

	return nil
}

func main() {
	SimpleAssetChaincodeContract := new(SimpleAssetChaincode)

	chaincode, err := contractapi.NewChaincode(SimpleAssetChaincodeContract)
	if err != nil {
		fmt.Printf("Error creating SimpleAsset chaincode: %s", err.Error())
		return	
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err.Error())
	}
}
