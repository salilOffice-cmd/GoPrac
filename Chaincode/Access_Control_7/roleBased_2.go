// In this lesson, we are going to learn about 'cid' library
// when using low level fabric api and also using high level fabric api

// Lesson Started==
// Write Chaincode for access control
// That is, give access to certain resources based on the function caller identity


// Verifying the signature on the client's certificate using the public key
// of the Certificate Authority (CA) involves cryptographic operations
// that typically occur outside of the chaincode itself. Hyperledger
// Fabric's SDKs and infrastructure handle certificate verification during
// transaction processing. However, within the chaincode, you can assume that
// the client's certificate has been authenticated by the network and focus
// on validating attributes or roles contained within it.

// Here's a brief overview of how the process works:

// 1. Client Authentication: When a client interacts with the blockchain network, 
// they present their X.509 certificate along with the transaction request.

// 2. Transaction Processing: The client's certificate, including its digital signature, 
// is included in the transaction payload.

// 3. Validation by Peers: Peers in the network validate the transaction, 
// which includes verifying the client's certificate against the CA's public key.

// 4. Chaincode Invocation: During chaincode execution, the chaincode has access to
// the client's certificate, which has already been authenticated by the peer.

// 5. Attribute Verification: The chaincode inspects the attributes or roles present
// in the client's certificate to determine access rights or permissions.

// 6. Access Control: Based on the verified attributes or roles,
// the chaincode enforces access control policies to allow or deny certain operations.


// USING LOW LEVEL APIs
// cid (used with low level fabric apis)
// The client identity(cid) chaincode library enables you to write
// chaincode which makes access control decisions based on the
// identity of the client (i.e. the invoker of the chaincode). 
// To read more on cid -->
// https://hyperledger-fabric.readthedocs.io/en/release-2.2/developapps/transactioncontext.html#clientidentity
//  Or 
// https://pkg.go.dev/github.com/hyperledger/fabric-chaincode-go/pkg/cid

// ClientIdentity takes the information returned by getCreator()
// and puts a set of X.509 utility APIs on top of it to make it easier to use for this common use case.



package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type DocumentChaincode struct{}

func (cc *DocumentChaincode) UploadDocument(stub shim.ChaincodeStubInterface, docName string) peer.Response {

	// Get the user's department attribute
	departmentBytes, err := cid.GetAttributeValue(stub, "department") // ****
	if err != nil {
		return shim.Error("failed to get user's department")
	}
	department := string(departmentBytes)


	// Get the user's role attribute
	roleBytes, err := cid.GetAttributeValue(stub, "role")
	if err != nil {
		return shim.Error("failed to get user's role")
	}
	role := string(roleBytes)


	// Check if the user is authorized to upload documents
	if role != "admin" && department != "IT" {
		return shim.Error("only IT admins can upload documents")
	}

	// Get user ID and MSP ID
	userIDBytes, err := cid.getID()
	if err != nil {
		return shim.Error(err.Error())
	}
	userID := string(userIDBytes)

	userMSPID, err := cid.GetMSPID()
	if err != nil {
		return shim.Error(err.Error())
	}


	// Log user ID and MSP ID
	fmt.Printf("User %s from MSP %s uploaded document '%s'\n", userID, userMSPID, docName)

	// Proceed with document upload logic...
	fmt.Printf("Document '%s' uploaded successfully\n", docName)
	return shim.Success([]byte("Document uploaded successfully"))


	// to read all the functions of cid -
	// https://pkg.go.dev/github.com/hyperledger/fabric-chaincode-go/pkg/cid#ClientIdentity
}


func main() {
	err := shim.Start(new(DocumentChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}



// USING HIGH LEVEL APIs
// Here the high level api internally calls the cid library
package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DocumentChaincode struct {
	contractapi.Contract
}

func (cc *DocumentChaincode) UploadDocument(ctx contractapi.TransactionContextInterface, docName string) error {
	// Get the user's department attribute
	department, err := ctx.GetClientIdentity().GetAttributeValue("department")
	if err != nil {
		return fmt.Errorf("failed to get user's department: %v", err)
	}

	// Get the user's role attribute
	role, err := ctx.GetClientIdentity().GetAttributeValue("role")
	if err != nil {
		return fmt.Errorf("failed to get user's role: %v", err)
	}

	// Check if the user is authorized to upload documents
	if role != "admin" && department != "IT" {
		return errors.New("only IT admins can upload documents")
	}

	// Get user ID and MSP ID
	userID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("failed to get user ID: %v", err)
	}

	userMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get user MSP ID: %v", err)
	}

	// Log user ID and MSP ID
	fmt.Printf("User %s from MSP %s uploaded document '%s'\n", userID, userMSPID, docName)

	// Proceed with document upload logic...
	fmt.Printf("Document '%s' uploaded successfully\n", docName)

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&DocumentChaincode{})
	if err != nil {
		fmt.Printf("Error creating DocumentChaincode: %s", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}


