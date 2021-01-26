package unmarshal

import (
	"fmt"
	"testing"
)

func TestContractUnmarshal(t *testing.T) {
	fmt.Println("testing contractunmarshal.go")
	t.Run("testing unmarshalcontract", TestUnmarshalContract)
}

func TestUnmarshalContract(t *testing.T) {
	contract := "Version: 1, Sequence: 1, Endorsement Plugin: escc, Validation Plugin: vscc, Approvals: [Org1MSP: true, Org2MSP: true],Name:fabcar,Channel:mychannel\n"
	contracts := UnmarshalContract([]byte(contract))
	for k, v := range contracts {
		fmt.Printf("No.%v===%v\n", k, v)
	}
}
