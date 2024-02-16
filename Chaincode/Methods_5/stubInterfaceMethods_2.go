// These are the actual functions that we will use the interact with ledger

// Functions available through the Chaincode Stub:

// GetArgs: Retrieves the arguments passed to the chaincode invocation.
// GetFunctionAndParameters: Parses the chaincode invocation to extract the function name and parameters.
// GetTxID: Retrieves the transaction ID of the current transaction.
// GetTxTimestamp: Retrieves the timestamp of the current transaction.
// GetCreator: Retrieves the identity of the client (creator) that initiated the transaction.
// GetChannelID: Retrieves the ID of the channel on which the transaction was submitted.
// PutState: Stores a key-value pair in the ledger's key-value store.
// GetState: Retrieves the value associated with a specified key from the ledger's key-value store.
// DelState: Deletes the value associated with a specified key from the ledger's key-value store.
// GetStateByRange: Retrieves a range of key-value pairs from the ledger's key-value store based on key range.
// GetStateByPartialCompositeKey: Retrieves key-value pairs from the ledger's key-value store based on a partial composite key.
// GetQueryResult: Executes a rich query against the state database and returns an iterator for the result set.
// InvokeChaincode: Invokes another chaincode on the same or a different channel.
// GetPrivateData: Retrieves private data associated with a specified key.
// PutPrivateData: Stores private data associated with a specified key.
// GetPrivateDataByRange: Retrieves a range of private data based on key range.
// GetPrivateDataByPartialCompositeKey: Retrieves private data based on a partial composite key.


// 1. GetArgs retrieves the arguments of the transaction.
func (cc *SimpleChaincode) GetArgsExample(stub shim.ChaincodeStubInterface) [][]byte {
	// The outer [] denotes a slice, meaning it can hold multiple elements of the same type.
	// The inner []byte denotes a byte slice, which is a sequence of bytes.
	// This type is often used to represent the arguments passed to a chaincode function
	args := stub.GetArgs()
	return args
}


// 2. GetFunctionAndParameters retrieves the function name and parameters of the transaction.
func (cc *SimpleChaincode) GetFunctionAndParametersExample(stub shim.ChaincodeStubInterface) (string, []string) {
	function, parameters := stub.GetFunctionAndParameters()
	return function, parameters
}


// 3. GetTxID retrieves the transaction ID of the current transaction.
func (cc *SimpleChaincode) GetTxIDExample(stub shim.ChaincodeStubInterface) string {
	txID := stub.GetTxID()
	return txID
}


// 4. GetTxTimestamp retrieves the timestamp of the current transaction.
func (cc *SimpleChaincode) GetTxTimestampExample(stub shim.ChaincodeStubInterface) (time.Time, error) {
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return time.Time{}, err
	}
	return timestamp, nil
}


// 5. GetCreatorExample retrieves the identity of the creator (submitter) of the transaction.
// (more about this in 7th lesson)
func (cc *SimpleChaincode) GetCreatorExample(stub shim.ChaincodeStubInterface) ([]byte, error) {
	creator, err := stub.GetCreator()
	if err != nil {
		return nil, err
	}
	return creator, nil
}


// 6. GetChannelIDExample retrieves the ID of the channel on which the transaction is executed.
func (cc *SimpleChaincode) GetChannelIDExample(stub shim.ChaincodeStubInterface) string {
	channelID := stub.GetChannelID()
	return channelID
}


// 7. PutStateExample writes a key-value pair to the world state.
func (cc *SimpleChaincode) PutStateExample(stub shim.ChaincodeStubInterface, key string, value []byte) error {
	err := stub.PutState(key, value)
	return err
}
