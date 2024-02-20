package main

import (
	"encoding/json"
	"fmt"

  	"github.com/hyperledger/fabric-chaincode-go/shim"
  	pb "github.com/hyperledger/fabric-protos-go/peer"

	// T
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

// SimpleAssetChaincode defines the Smart Contract structure
type SimpleAssetChaincode struct{}

// Asset represents a single asset
type Asset struct {
	ID    string `json:"ID"`
	Owner string `json:"owner"`
	Color string `json:"color"`
	Size  int    `json:"size"`
	Price int    `json:"price"`
}

// INIT AND INVOKE

// INIT (can be considered as a contructor for the contract)
// When using the low level api for Go, both the Init and Invoke functions are essential.
func (s *SimpleAssetChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// Initialize chaincode
    // The Init function is invoked when the chaincode is instantiated or upgraded.
    // It's typically used to perform any necessary initialization tasks, such as setting up
    // initial state or configuration.
	return shim.Success(nil)
}

// INVOKE (mediator for the external application and the custom functions of the contract)
func (s *SimpleAssetChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// The Invoke function is the entry point for processing transactions in the chaincode.
    // When a transaction is submitted to the chaincode, the Invoke function is called.
    // Inside the Invoke function, the function name and parameters are extracted using
    // ctx.GetStub().GetFunctionAndParameters().
    // Then, based on the extracted function name, the transaction is routed
    // to the appropriate handler function (e.g., CreateAsset, UpdateAsset, QueryAsset).
	
	function, args := stub.GetFunctionAndParameters()

	// Route to the appropriate handler function
	if function == "CreateAsset" {
		return s.CreateAsset(stub, args)
	} else if function == "UpdateAsset" {
		return s.UpdateAsset(stub, args)
	} else if function == "QueryAsset" {
    return s.QueryAsset(stub, args)
  }

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SimpleAssetChaincode) CreateAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var asset Asset
	err := json.Unmarshal([]byte(args[0]), &asset)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal asset: %s", err))
	}

	exists, err := s.AssetExists(stub, asset.ID)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to check asset existence: %s", err))
	}
	if exists {
		return shim.Error(fmt.Sprintf("Asset %s already exists", asset.ID))
	}

	err = stub.PutState(asset.ID, []byte(args[0]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", err))
	}

	return shim.Success(nil)
}

func (s *SimpleAssetChaincode) AssetExists(stub shim.ChaincodeStubInterface, assetID string) (bool, error) {
	assetBytes, err := stub.GetState(assetID)
	if err != nil {
		return false, fmt.Errorf("Failed to read from world state: %v", err)
	}
	return assetBytes != nil, nil
}

func (s *SimpleAssetChaincode) UpdateAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var asset Asset
	err := json.Unmarshal([]byte(args[0]), &asset)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal asset: %s", err))
	}

	exists, err := s.AssetExists(stub, asset.ID)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to check asset existence: %s", err))
	}
	if !exists {
		return shim.Error(fmt.Sprintf("Asset %s does not exist", asset.ID))
	}

	err = stub.PutState(asset.ID, []byte(args[0]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update asset: %s", err))
	}

	return shim.Success(nil)
}

func (s *SimpleAssetChaincode) QueryAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assetBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to read asset %s from world state: %v", args[0], err))
	}
	if assetBytes == nil {
		return shim.Error(fmt.Sprintf("Asset %s does not exist", args[0]))
	}

	return shim.Success(assetBytes)
}

func main() {
	// err := shim.Start(new(SimpleAssetChaincode))
	// if err != nil {
	//   fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	// }
	testChainCode()
}


// TESTING
func testChainCode() {
	cc := new(SimpleAssetChaincode)

	// Create a new shim test instance with the mock chaincode
	stub := shimtest.NewMockStub("TestStub", cc)

	// Define the asset to be created
	asset := Asset{
		ID:    "asset3",
		Owner: "Charlie",
		Color: "green",
		Size:  8,
		Price: 300,
	}

	// Convert asset to bytes
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		fmt.Printf("Error marshalling asset: %s", err)
		return
	}

	// Invoke the CreateAsset function
	response := stub.MockInvoke("CreateAsset", [][]byte{[]byte("CreateAsset"), assetBytes})

	// Check if the invocation was successful
	if response.Status != shim.OK {
		fmt.Printf("Error invoking chaincode: %s", response.Message)
		return
	}

  if response.Status == shim.OK {
    fmt.Println("CreateAsset executed successfully")
  }

	// Query the asset to verify if it was created
	queryResponse := stub.MockInvoke("QueryAsset", [][]byte{[]byte("QueryAsset"), []byte(asset.ID)})
	if queryResponse.Status != shim.OK {
		fmt.Printf("Error querying asset: %s", queryResponse.Message)
		return
	}

	// Unmarshal the asset returned from the query
	var queriedAsset Asset
	err = json.Unmarshal(queryResponse.Payload, &queriedAsset)
	if err != nil {
		fmt.Printf("Error unmarshalling asset: %s", err)
		return
	}

	// Check if the queried asset matches the original asset
	if queriedAsset.ID != asset.ID || queriedAsset.Owner != asset.Owner || queriedAsset.Color != asset.Color || queriedAsset.Size != asset.Size || queriedAsset.Price != asset.Price {
		fmt.Println("Queried asset does not match original asset")
		return
	}

	fmt.Println("Test passed: Asset creation and querying successful")
}



