// DEFER KEYWORD IN GO
// Defer in english means postpone
// In Go, defer is a keyword used to schedule a function call to be executed
// when the surrounding function returns. It is commonly used to ensure that
// resources are released or cleanup actions are performed regardless of whether
// the function exits normally or panics.

// Here's an example of how defer can be used in a chaincode function:
// Using defer in normal functions
func (s *SimpleAssetChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, asset Asset) error {
    // Open a database connection
    db, err := OpenDatabase()
    if err != nil {
        return err
    }
    defer db.Close() // Schedule the database connection to be closed when the function returns

    // Perform database operations
    err = db.Insert(asset)
    if err != nil {
        return err
    }

    // If no errors occurred, return nil
    return nil
}





// USING DEFER IN INVOKE() FUNCTION
// When a panic exception occurs, the chaincode container may be suspended and
// restarted, logs cannot be found, and the problem cannot be located immediately.
// To prevent this case, add the defer statement at the entry point of the Invoke
// function. When a panic occurs, the error is returned to the client.
func (s *SimpleAssetChaincode) Invoke(stub shim.ChaincodeStubInterface) (string, error) {
    

     defer func() {
        if err := recover(); err != nil {
            // Handle the panic by returning an error response  
            errMsg := fmt.Sprintf("Chaincode panicked: %v", err)
            logger.Errorf(errMsg)
            panic(errMsg)
        }
    }()
    
    // Actual Invoke logic goes here...
    function, args := ctx.GetStub().GetFunctionAndParameters()

    // Route to the appropriate handler function
    if function == "CreateAsset" {
        return s.CreateAsset(ctx, args)
    } else if function == "UpdateAsset" {
        return s.UpdateAsset(ctx, args)
    } else if function == "QueryAsset" {
        return s.QueryAsset(ctx, args)
    }

    return nil, fmt.Errorf("Invalid Smart Contract function name.")
}


// FLOW OF DEFER IN INVOKE()

// defer Statement:
// 1. This defer statement defines an anonymous function that will be executed
// when the surrounding function (Invoke) exits. It's deferred until the end of the Invoke function.
// 2. Inside the deferred function, recover() is called to capture any panic
// that occurs within the surrounding function.
// 3. If a panic is recovered (i.e., if recover() returns a non-nil value),
// the deferred function handles the panic by logging an error message,
// constructing an error response, and then panicking again with the error response encoded as JSON.


// Invoke Function Flow:

// 1. The Invoke function starts by deferring the error recovery logic to ensure that any panics during its execution are captured and handled.
// 2. It then proceeds to extract the function name and arguments from the chaincode stub using GetFunctionAndParameters().
// 3. Based on the extracted function name (function), it routes the invocation to the corresponding function (invoke, delete, or query).
// 4. If the function name is "delete", the control will pass to the delete function.
// 5. However, if a panic occurs during the execution of the delete function (or any other function invoked within Invoke), the deferred error recovery logic will be triggered.
// 6. In case of a panic, the deferred function logs an error, constructs an error response with status code 500 and an error message, marshals it into JSON, and then panics again with the JSON-encoded error response.
// 7. This final panic will bubble up and terminate the chaincode execution, but it ensures that the panic is handled gracefully, preventing the chaincode container from crashing and providing an informative error response to the caller.


// Panic() function

// In Go, panic is a built-in function that stops the ordinary flow of control
// and begins panicking. When a panic occurs, the function that is currently executing
// is terminated, any deferred functions are executed, and then the panic is propagated 
// up the call stack until it reaches the top-level function of the goroutine. 
// If that goroutine was invoked by the runtime, the program will terminate with a non-zero exit code.

// However, it's worth noting that using panic for error handling is generally
// considered a last resort and should be used sparingly. 
// It's typically preferable to handle errors explicitly using error values
// and appropriate control flow mechanisms. Panics are intended for unrecoverable errors,
// such as programming errors or invalid states, rather than recoverable errors that
// can be anticipated and managed.
