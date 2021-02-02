package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// Chaincode implementation
type ChainCode struct {
}

type TransactionInfo struct {
	TransactionID  string       `json:transactionID`
	ProductName    string       `json:productName`
	ProductPrice   string       `json:productPrice`
	ProductDes     string       `json:productDes`
	ServicePeriod  string       `json:servicePeriod`
	TransactionPro ProviderInfo `json:transactionPro`
	TransactionBuy BuyerInfo    `json:transactionBuy`
}

type TransactionOnly struct {
	TransactionOnlyID   string `json:transactionOnlyID`
	TransactionOnlyLoan string `json:transactionOnlyLoan`
	TransactionOnlyTime string `json:transactionOnlyTime`
	TransactionOnlyRate string `json:transactionOnlyRate`
}

type ProviderInfo struct {
	ProviderName   string `json:providerName`
	ProviderLedger string `json:providerLedger`
	ProviderCredit string `json:providerCredit`
}

type BuyerInfo struct {
	BuyerName   string `json:buyerName`
	BuyerLedger string `json:buyerLedger`
	BuyerCredit string `json:buyerCredit`
}

func (t *ChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	var transactionInfos TransactionInfo

	transactionInfos.TransactionID = "20060102150405"                      //交易ID
	transactionInfos.ProductName = "源码加固"                                  //产品名称
	transactionInfos.ProductPrice = "100000"                               //产品价格
	transactionInfos.ProductDes = "支持在线，支持C/C++/OC语言开发的iOS、Android应用源代码加固" //产品描述
	transactionInfos.ServicePeriod = "1"                                   //服务周期
	transactionInfos.TransactionPro.ProviderName = "aaaa"                  //提供商名称
	transactionInfos.TransactionPro.ProviderCredit = "A"                   //提供商信用
	transactionInfos.TransactionPro.ProviderLedger = "50000"               //提供商账户余额
	transactionInfos.TransactionBuy.BuyerName = "梆梆安全"                     //购买方名称
	transactionInfos.TransactionBuy.BuyerCredit = "A"                      //购买方信用
	transactionInfos.TransactionBuy.BuyerLedger = "500000"                 //购买方账户余额
	transactionInfosJSONasBytes, err := json.Marshal(transactionInfos)
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	err = stub.PutState(transactionInfos.TransactionID, transactionInfosJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *ChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "querypromoney" {
		// Deletes an entity from its state
		return t.querypromoney(stub, args)
	} else if function == "querybuymoney" {
		// the old "Query" is now implemtned in invoke
		return t.querybuymoney(stub, args)
	} else if function == "querytransaction" {
		// the old "Query" is now implemtned in invoke
		return t.querytransaction(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *ChainCode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var providerledger, buyerledger int // Asset holdings
	transactionInfos := new(TransactionInfo)
	var X int // Transaction value
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//transactionInfos.TransactionID = args[0]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	tranBytes, err := stub.GetState(args[0])
	fmt.Printf("=======+%v\n", tranBytes)

	//resultsIterator, err := stub.GetHistoryForKey(transactionInfos.TransactionID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	//defer resultsIterator.Close()

	var transactionPro ProviderInfo
	var transactionBuy BuyerInfo
	json.Unmarshal(tranBytes, transactionInfos)
	if (transactionInfos.TransactionPro.ProviderLedger != "") && (transactionInfos.TransactionBuy.BuyerLedger != "") {
		transactionPro = transactionInfos.TransactionPro
		transactionBuy = transactionInfos.TransactionBuy
	}
	providerledger, _ = strconv.Atoi(string(transactionPro.ProviderLedger))
	buyerledger, _ = strconv.Atoi(string(transactionBuy.BuyerLedger))

	X, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	providerledger = providerledger + X
	buyerledger = buyerledger - X
	fmt.Printf("providerledger = %d, buyerledger = %d\n", providerledger, buyerledger)

	transactionInfos.TransactionPro.ProviderLedger = strconv.Itoa(providerledger)
	transactionInfos.TransactionBuy.BuyerLedger = strconv.Itoa(buyerledger)
	// Write the state back to the ledger
	transactionInfosJSONasBytes, err := json.Marshal(transactionInfos)
	fmt.Printf("=====%+v\n", transactionInfos)
	fmt.Printf("=====%v\n", transactionInfosJSONasBytes)

	err = stub.PutState(args[0], transactionInfosJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *ChainCode) querypromoney(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	transactionInfos := new(TransactionInfo)
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//transactionInfos.TransactionID = args[0]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	//resultsIterator, err := stub.GetHistoryForKey(transactionInfos.TransactionID)
	tranBytes, err := stub.GetState(args[0])
	fmt.Printf("=======+%v\n", tranBytes)

	if err != nil {
		return shim.Error("Failed to get state")
	}
	//defer resultsIterator.Close()

	var transactionPro ProviderInfo
	json.Unmarshal(tranBytes, transactionInfos)
	fmt.Printf("=======+%v\n", transactionInfos)
	if (transactionInfos.TransactionPro.ProviderLedger != "") && (transactionInfos.TransactionBuy.BuyerLedger != "") {
		transactionPro = transactionInfos.TransactionPro
	}
	fmt.Printf("=======+%v\n", transactionPro)
	jsonsAsBytes, err := json.Marshal(transactionPro)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Query Response:%s\n", jsonsAsBytes)
	return shim.Success(jsonsAsBytes)
	//return transactionPro
}

func (t *ChainCode) querybuymoney(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	transactionInfos := new(TransactionInfo)
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//transactionInfos.TransactionID = args[0]

	tranBytes, err := stub.GetState(args[0])
	fmt.Printf("=======+%v\n", tranBytes)

	if err != nil {
		return shim.Error("Failed to get state")
	}
	//defer resultsIterator.Close()

	var transactionBuy BuyerInfo
	json.Unmarshal(tranBytes, &transactionInfos)
	fmt.Printf("=======+%v\n", transactionInfos)
	if (transactionInfos.TransactionPro.ProviderLedger != "") && (transactionInfos.TransactionBuy.BuyerLedger != "") {
		transactionBuy = transactionInfos.TransactionBuy
	}

	jsonsAsBytes, err := json.Marshal(transactionBuy)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func (t *ChainCode) querytransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionInfos TransactionInfo
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	transactionInfos.TransactionID = args[0]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	resultsIterator, err := stub.GetHistoryForKey(transactionInfos.TransactionID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &transactionInfos)

	}

	jsonsAsBytes, err := json.Marshal(transactionInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func main() {
	shim.Start(new(ChainCode))
}
