// Define the chaincode
type SimpleAssetChaincode struct {
    contractapi.Contract
}

// Function to create a new asset
func (s *SimpleAssetChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, assetID string, owner string) error {
    // Create the asset
    asset := Asset{
        ID:    assetID,
        Owner: owner,
    }

    // Store the asset in the world state
    err := ctx.GetStub().PutState(assetID, []byte(owner))
    if err != nil {
        return err
    }

    // Emit an event indicating a new asset was created
    err = ctx.GetStub().SetEvent("AssetCreated", []byte(assetID))
    if err != nil {
        return fmt.Errorf("error emitting event: %v", err)
    }

    return nil
}
