// A smart contract can access a range of functionality in a smart contract
// via the transaction context 'stub' and 'clientIdentity'.

// In this lesson, we have performed CRUD operations using 'STUB'


// To get all the functions of GetStub().{function}
// refer 
// https://hyperledger-fabric.readthedocs.io/en/release-2.2/developapps/transactioncontext.html#stub
// OR
// https://pkg.go.dev/github.com/hyperledger/fabric-chaincode-go/shim#ChaincodeStubInterface

// 1. Creating Asset: Use the PutState method of the stub to store the new asset in the ledger.
func (s *SimpleAssetChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, asset Asset) error {
    // Check if asset already exists
    exists, err := s.AssetExists(ctx, asset.ID)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the asset %s already exists", asset.ID)
    }

    // Serialize asset to JSON
    // This step is done as the PutState() function that stores key value pair in the ledger
    // accepts the value in json format and not in go data types
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return fmt.Errorf("failed to marshal asset to JSON: %v", err)
    }

    // Put asset in the ledger
    err = ctx.GetStub().PutState(asset.ID, assetJSON)
    if err != nil {
        return fmt.Errorf("failed to put asset in ledger: %v", err)
    }

    return nil
}



// 2. UpdateAsset: Use the PutState method of the stub to update the existing asset in the ledger.
func (s *SimpleAssetChaincode) UpdateAsset(ctx contractapi.TransactionContextInterface, updatedAsset Asset) error {
   
    // To update an asset, logically it makes sense that we should check if the key exists in ledger or not
    // Check if asset already exists
    exists, err := s.AssetExists(ctx, updatedAsset.ID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exists", asset.ID)
    }


    // Serialize updated asset to JSON
    updatedAssetJSON, err := json.Marshal(updatedAsset)
    if err != nil {
        return fmt.Errorf("failed to marshal updated asset to JSON: %v", err)
    }

    // Update asset in the ledger
    err = ctx.GetStub().PutState(updatedAsset.ID, updatedAssetJSON)
    if err != nil {
        return fmt.Errorf("failed to update asset in ledger: %v", err)
    }

    return nil
}


// 3. TransferAssetOwnership: Use the PutState method of the stub to update the ownership field of the existing asset in the ledger.
func (s *SimpleAssetChaincode) TransferAssetOwnership(ctx contractapi.TransactionContextInterface, assetID string, newOwner string) error {
    
    // Retrieve existing asset from the ledger
    assetJSON, err := ctx.GetStub().GetState(assetID)
    if err != nil {
        return fmt.Errorf("failed to read asset from ledger: %v", err)
    }
    if assetJSON == nil {
        return fmt.Errorf("asset %s does not exist", assetID)
    }

    // Deserialize existing asset JSON
    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return fmt.Errorf("failed to unmarshal asset JSON: %v", err)
    }

    // Update ownership field
    asset.Owner = newOwner

    // Serialize updated asset to JSON
    updatedAssetJSON, err := json.Marshal(asset)
    if err != nil {
        return fmt.Errorf("failed to marshal updated asset to JSON: %v", err)
    }

    // Update asset in the ledger
    err = ctx.GetStub().PutState(assetID, updatedAssetJSON)
    if err != nil {
        return fmt.Errorf("failed to update asset in ledger: %v", err)
    }

    return nil
}


// 4. DeleteAsset: Use the DelState method of the stub to delete the existing asset from the ledger.
func (s *SimpleAssetChaincode) DeleteAsset(ctx contractapi.TransactionContextInterface, assetID string) error {

    // Check if asset exists
    exists, err := s.AssetExists(ctx, updatedAsset.ID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exists", asset.ID)
    }

    // Delete asset from the ledger
    err := ctx.GetStub().DelState(assetID)
    if err != nil {
        return fmt.Errorf("failed to delete asset from ledger: %v", err)
    }

    return nil
}


// 5. AssetExists: Use the GetState method of the stub to check if an asset exists in the ledger.
func (s *SimpleAssetChaincode) AssetExists(ctx contractapi.TransactionContextInterface, assetID string) (bool, error) {
    // Check if asset exists in the ledger
    assetJSON, err := ctx.GetStub().GetState(assetID)
    if err != nil {
        return false, fmt.Errorf("failed to read asset from ledger: %v", err)
    }
    return assetJSON != nil, nil
}


// 6. QueryAssetByID: Use the GetState method of the stub to retrieve the asset from the ledger by its ID.
func (s *SimpleAssetChaincode) QueryAssetByID(ctx contractapi.TransactionContextInterface, assetID string) (Asset, error) {
    // Retrieve asset from the ledger
    assetJSON, err := ctx.GetStub().GetState(assetID)
    if err != nil {
        return nil, fmt.Errorf("failed to read asset from ledger: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("asset %s does not exist", assetID)
    }

    // Deserialize asset JSON
    //(only required when you want to use this asset anywhere in this method, not in this example)
    //var asset Asset
    //err = json.Unmarshal(assetJSON, &asset)
    //if err != nil {
    //    return nil, fmt.Errorf("failed to unmarshal asset JSON: %v", err)
    //}
    //return asset, nil

    return assetJSON, nil
}


// 7.QueryAllAssets: Use the GetStateByRange method of the stub to retrieve all assets from the ledger.
func (s *SimpleAssetChaincode) QueryAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
    // Retrieve all assets from the ledger
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve assets from ledger: %v", err)
    }
    defer resultsIterator.Close()

    var assets []*Asset
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, fmt.Errorf("failed to iterate over results: %v", err)
        }

        var asset Asset
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, fmt.Errorf("failed to unmarshal asset JSON: %v", err)
        }
        assets = append(assets, &asset) 
	// the above code is written if we want to modify or use all these assets
	
	// But when we dont want to use this assets list anywhere in this method, we can simply return it 
	// like this -->
	// var assets []Asset
	// assets = append(assets, asset);
	// return assets, nil
    }

    return assets, nil



    // Learnings from this method:

    // resultsIterator:
    // The resultsIterator returned by GetStateByRange is indeed an iterator, not a storage variable.
    // It does not store all key-value pairs in memory at once.
    // Instead, it provides a mechanism to iterate over the key-value pairs one by one, 
    // fetching them from the database as needed. This approach is memory-efficient and 
    // allows you to process large datasets without consuming excessive memory.

    // resultsIterator.Next():
    // When you use resultsIterator.Next() in a loop, each call to Next() fetches the 
    // next key-value pair from the database. The iterator fetches key-value pairs lazily,
    // meaning it retrieves them on-demand as you iterate over them. 
    // This can help reduce the memory footprint of your chaincode, especially when dealing with large datasets.
}


// 8. GetHistoryForAsset: Use the GetHistoryForKey method of the stub to retrieve the transaction history for the asset by its ID.
func (s *SimpleAssetChaincode) GetHistoryForAsset(ctx contractapi.TransactionContextInterface, assetID string) ([]*TransactionHistory, error) {
    // Retrieve transaction history for the asset from the ledger
    resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetID)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve history for asset from ledger: %v", err)
    }
    defer resultsIterator.Close()

    var history []*TransactionHistory
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, fmt.Errorf("failed to iterate over results: %v", err)
        }

        // Construct transaction history object
        transaction := TransactionHistory{
            TxId:      queryResponse.TxId,
            Value:     queryResponse.Value,
            Timestamp: time.Unix(queryResponse.Timestamp.Seconds, int64(queryResponse.Timestamp.Nanos)),
        }
        history = append(history, &transaction)
    }

    return history, nil
}

TransactionHistory struct might look something like this:
type TransactionHistory struct {
    TxId      string    `json:"txId"`
    Value     []byte    `json:"value"`
    Timestamp time.Time `json:"timestamp"`
}
