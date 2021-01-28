package unmarshal

import (
	"bytes"
	"fmt"
)

type Contract struct {
	Name       string
	Channel    string
	Version    string
	Endorse    string
	Validation string
	Approve    string
}

func UnmarshalContract(contract []byte) []Contract {
	if len(contract) != 0 {
		var contracts []Contract
		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(": ,[]"), r)
		}
		contractBytes := bytes.FieldsFunc(contract, f)
		fmt.Printf("Fields are: %q\n", contractBytes)
		for i := 0; i < len(contractBytes)/19; i++ {
			var contractFields Contract
			contractBytes[18+19*i] = bytes.Replace(contractBytes[18+19*i], []byte("\n"), []byte(""), -1)
			fmt.Printf("replace are: %q\n", contractBytes)
			contractFields.Version = string(contractBytes[1+19*i])
			contractFields.Endorse = string(contractBytes[6+19*i])
			contractFields.Validation = string(contractBytes[9+19*i])
			contractFields.Name = string(contractBytes[16+19*i])
			contractFields.Channel = string(contractBytes[18+19*i])
			contractFields.Approve = string(contractBytes[11+19*i]) + ": " + string(contractBytes[12+19*i]) + ", " + string(contractBytes[13+19*i]) + ": " + string(contractBytes[14+19*i])
			contracts = append(contracts, contractFields)
		}
		return contracts
	} else {
		return nil
	}
}
