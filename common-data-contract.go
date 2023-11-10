package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const COMMONDATA = "commondata"

// CommonDataContract contract for handling CommonContractData
type CommonDataContract struct {
	contractapi.Contract
}

// NewCommonData adds a new CommonData to the world state using id as key
func (cc *CommonDataContract) NewCommonData(ctx CustomTransactionContextInterface, id string, value string) error {
	existing := ctx.GetData()

	if existing != nil {
		return fmt.Errorf("Cannot create new CommonContractData in world state as key %s already exists", id)
	}

	commonData := &CommonData{
		DocType: COMMONDATA,
		ID:      id,
		Value:   value,
	}
	cdBytes, err := json.Marshal(commonData)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(id, cdBytes)
	if err != nil {
		return err
	}
	cde := &CommonDataEvent{
		Key:   commonData.ID,
		Value: commonData.Value,
	}
	cdeBytes, _ := json.Marshal(cde)
	err = SendEvent(ctx, "NewCommonDataEvent", cdeBytes)
	if err != nil {
		return fmt.Errorf("call SendEvent failed:%s", err.Error())
	}
	return nil
}

//UpdateCommonData  更新通用数据
func (cc *CommonDataContract) UpdateCommonData(ctx CustomTransactionContextInterface, id string, newValue string) error {
	existing := ctx.GetData()

	if existing == nil {
		return fmt.Errorf("Cannot update CommonData in world state as key %s does not exist", id)
	}

	cd := new(CommonData)

	err := json.Unmarshal(existing, cd)

	if err != nil {
		return fmt.Errorf("Data retrieved from world state for key %s was not of type CommonData", id)
	}

	cd.Value = newValue

	cdBytes, _ := json.Marshal(cd)

	err = ctx.GetStub().PutState(id, []byte(cdBytes))

	if err != nil {
		return fmt.Errorf("Unable to interact with world state")
	}

	cde := &CommonDataEvent{
		Key:   cd.ID,
		Value: cd.Value,
	}
	cdeBytes, _ := json.Marshal(cde)

	err = SendEvent(ctx, "UpdateCommonDataEvent", cdeBytes)
	if err != nil {
		return fmt.Errorf("call SendEvent failed:%s", err.Error())
	}

	return nil
}

// DeleteCommonData delete a CommonData with given id from world state
func (cc *CommonDataContract) DeleteCommonData(ctx CustomTransactionContextInterface, id string) error {
	existing := ctx.GetData()

	if existing == nil {
		return fmt.Errorf("Cannot read world state pair with key %s. Does not exist", id)
	}
	cd := new(CommonData)

	err := json.Unmarshal(existing, cd)

	if err != nil {
		return fmt.Errorf("Data retrieved from world state for key %s was not of type CommonData", id)
	}
	err = ctx.GetStub().DelState(id) //remove the CommonData from chaincode state
	if err != nil {
		return fmt.Errorf("Delete a CommonData from world state for key %s failed", id)
	}

	cde := &CommonDataEvent{
		Key:   cd.ID,
		Value: cd.Value,
	}
	cdeBytes, _ := json.Marshal(cde)
	err = SendEvent(ctx, "DeleteCommonDataEvent", cdeBytes)
	if err != nil {
		return fmt.Errorf("call SendEvent failed:%s", err.Error())
	}

	return nil
}

// QueryCommonData returns the organization with id given from the world state
func (cc *CommonDataContract) QueryCommonData(ctx CustomTransactionContextInterface, id string) (*CommonData, error) {
	existing := ctx.GetData()

	if existing == nil {
		return nil, fmt.Errorf("Cannot read world state pair with key %s. Does not exist", id)
	}

	cd := new(CommonData)

	err := json.Unmarshal(existing, cd)

	if err != nil {
		return nil, fmt.Errorf("Data retrieved from world state for key %s was not of type CommonData", id)
	}

	return cd, nil
}

// GetEvaluateTransactions returns functions of BroadcastContract not to be tagged as submit
func (cc *CommonDataContract) GetEvaluateTransactions() []string {
	return []string{"QueryCommonData"}
}
