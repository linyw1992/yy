package main

//CommonData 通用增删改查合约数据
type CommonData struct {
	DocType string `json:"docType"`
	ID      string `json:"id"`
	Value   string `json:"value"`
}

//CommonDataEvent 通用增删改查合约数据合约
type CommonDataEvent struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
