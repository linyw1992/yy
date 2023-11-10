package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"os"
)

func main() {

	commonDataContract := new(CommonDataContract)
	commonDataContract.TransactionContextHandler = new(CustomTransactionContext)
	commonDataContract.BeforeTransaction = GetWorldState
	commonDataContract.AfterTransaction = SendTxEvent
	commonDataContract.Name = "commonOrg.CommonDataContract"
	commonDataContract.UnknownTransaction = UnknownTransactionHandler

	cc, err := contractapi.NewChaincode(commonDataContract)
	cc.DefaultContract = commonDataContract.GetName()

	if err != nil {
		panic(err.Error())
	}

	server := &shim.ChaincodeServer{
		CCID:    os.Getenv("CHAINCODE_CCID"),
		Address: os.Getenv("CHAINCODE_ADDRESS"),
		CC:      cc,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	// Start the chaincode external server
	err = server.Start()
	if err != nil {
		panic(err.Error())
	}
}
