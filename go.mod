module github.com/hyperledger/fabric-samples/chaincode/contracts/common-contracts

go 1.14

require (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20201119163726-f8ef75b17719 // indirect
	github.com/hyperledger/fabric-contract-api-go v1.1.1
	github.com/tjfoc/gmsm v1.4.0 // indirect
)

replace (
	github.com/hyperledger/fabric-chaincode-go => ../fabric-chaincode-go
	github.com/tjfoc/gmsm => ../gmsm
)