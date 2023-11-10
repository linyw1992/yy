package main

import "fmt"

//SendEvent 发送交易事件
func SendEvent(ctx CustomTransactionContextInterface, eventName string, eventPayload []byte) error {
	err := ctx.GetStub().SetEvent(eventName, eventPayload)
	if err != nil {
		return fmt.Errorf("SetEvent %s failure:%s", eventName, err.Error())
	}

	return nil
}
