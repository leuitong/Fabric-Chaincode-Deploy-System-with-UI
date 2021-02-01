package unmarshal

import (
	"bytes"
	"fmt"
)

type Transaction struct {
	TransactionID string
	ProductName   string
	ProductPrice  string
	ProductDes    string
	ServicePeriod string
}

func UnmarshalTransaction(transaction []byte) Transaction {
	var transaction1 Transaction
	if len(transaction) != 0 {
		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(":,{}\""), r)
		}
		transactionBytes := bytes.FieldsFunc(transaction, f)
		fmt.Printf("Fields are: %q\n", transactionBytes)
		transaction1.TransactionID = string(transactionBytes[1])
		transaction1.ProductName = string(transactionBytes[3])
		transaction1.ProductPrice = string(transactionBytes[5])
		transaction1.ProductDes = string(transactionBytes[7])
		transaction1.ServicePeriod = string(transactionBytes[9])
		fmt.Printf("provider are: %v\n", transaction1)
		return transaction1
	} else {
		return transaction1
	}
}
