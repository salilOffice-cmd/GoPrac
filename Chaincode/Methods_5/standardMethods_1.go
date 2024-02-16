// In this lesson, we are looking at some custom(created by us) chaincode standard methods which are
//  used with 's' as shown below:
// func (s *SimpleAssetChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {}
// 
// Note the name of this methods are just naming conventions, you can change them as per your need.
// Also note that these functions are called from external applications.
// These methods does not directly interact with the ledger but uses the 'chaincode stub' inside their implementation
// and the stub thereafter interacts with ledger

// List of Functions:
// InitLedger: Initializes the ledger with initial data when the chaincode is instantiated or upgraded.
// CreateAsset: Creates a new asset in the ledger.
// UpdateAsset: Updates an existing asset in the ledger.
// TransferAssetOwnership: Transfers ownership of an asset to a new owner.
// DeleteAsset: Deletes an existing asset from the ledger.
// QueryAssetByID: Retrieves information about a specific asset from the ledger using its ID.
// QueryAllAssets: Retrieves information about all assets stored in the ledger.
// AssetExists: Checks if a specific asset exists in the ledger.
// GetHistoryForAsset: Retrieves the transaction history for a specific asset.


// 1. CreateAsset:

// Create a new asset
asset := Asset{
    ID:    "1",
    Owner: "Alice",
    Color: "Red",
    Size:  10,
    Price: 100,
}

// Call the CreateAsset function
err := s.CreateAsset(ctx, asset)
if err != nil {
    fmt.Printf("Error creating asset: %v\n", err)
    return
}


// 2. UpdateAsset:
// Update an existing asset
updatedAsset := Asset{
    ID:    "1",
    Owner: "Bob",
    Color: "Blue",
    Size:  20,
    Price: 200,
}

// Call the UpdateAsset function
err := s.UpdateAsset(ctx, updatedAsset)
if err != nil {
    fmt.Printf("Error updating asset: %v\n", err)
    return
}


// 3. TransferAssetOwnership:
// Transfer ownership of an asset
newOwner := "Charlie"

// Call the TransferAssetOwnership function
err := s.TransferAssetOwnership(ctx, "1", newOwner)
// "1": This parameter represents the unique identifier (ID) of the asset whose ownership is being transferred
if err != nil {
    fmt.Printf("Error transferring asset ownership: %v\n", err)
    return
}


// 4. DeleteAsset:
// Delete an existing asset
assetIDToDelete := "1"

// Call the DeleteAsset function
err := s.DeleteAsset(ctx, assetIDToDelete)
if err != nil {
    fmt.Printf("Error deleting asset: %v\n", err)
    return
}


// 5. QueryAssetByID:
// Query an asset by its ID
assetIDToQuery := "1"

// Call the QueryAssetByID function
asset, err := s.QueryAssetByID(ctx, assetIDToQuery)
if err != nil {
    fmt.Printf("Error querying asset by ID: %v\n", err)
    return
}

fmt.Println("Asset:", asset)


// 6. QueryAllAssets:
// Query all assets
assets, err := s.QueryAllAssets(ctx)
if err != nil {
    fmt.Printf("Error querying all assets: %v\n", err)
    return
}

fmt.Println("All assets:", assets)


// 7.AssetExists:
// Check if an asset exists
assetIDToCheck := "1"

// Call the AssetExists function
exists, err := s.AssetExists(ctx, assetIDToCheck)
if err != nil {
    fmt.Printf("Error checking asset existence: %v\n", err)
    return
}

fmt.Println("Asset exists:", exists)


// 8. GetHistoryForAsset:
// Get transaction history for an asset
assetIDToQuery := "1"

// Call the GetHistoryForAsset function
history, err := s.GetHistoryForAsset(ctx, assetIDToQuery)
if err != nil {
    fmt.Printf("Error getting history for asset: %v\n", err)
    return
}

fmt.Println("History for asset:", history)