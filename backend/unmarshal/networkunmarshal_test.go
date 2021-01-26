package unmarshal

import (
	"fmt"
	"testing"
)

func TestNetworkUnmarshal(t *testing.T) {
	fmt.Println("testing networkunmarshal.go")
	t.Run("testing unmarshalnetwork", TestUnmarshalNetwork)
}

func TestUnmarshalNetwork(t *testing.T) {
	//network := []byte{}
	network := []byte("CONTAINER ID        IMAGE            COMMAND             CREATED             STATUS              PORTS                              NAMES 384aad68d679        hyperledger/fabric-peer:latest      `peer node start`   25 seconds ago      Up 19 seconds       0.0.0.0:7051->7051/tcp   peer0.org1.example.com 01bcfa955ab1        hyperledger/fabric-peer:latest      `peer node start`   25 seconds ago      Up 18 seconds       7051/tcp, 0.0.0.0:9051->9051/tcp   peer0.org2.example.com 1321f4d75ec5        hyperledger/fabric-orderer:latest   `orderer`           25 seconds ago      Up 20 seconds       0.0.0.0:7050->7050/tcp             orderer.example.com")
	nodes := UnmarshalNetwork(network)
	fmt.Println(nodes)

	for k, v := range nodes {
		fmt.Printf("No.%v===%v\n", k, v)
	}
}
