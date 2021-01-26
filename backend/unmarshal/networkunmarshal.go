package unmarshal

import (
	"bytes"
	"fmt"
)

type Node struct {
	ID     string
	Image  string
	Status string
	Ports  string
	Names  string
}

func UnmarshalNetwork(network []byte) []Node {
	if len(network) != 0 {
		var nodes []Node
		//8 21 35
		networkBytes := bytes.Fields(network)
		fmt.Printf("Fields are: %q\n", networkBytes)
		var peer1, peer2, orderer Node
		peer1.ID = string(networkBytes[0])
		peer1.Image = string(networkBytes[1])
		peer1.Status = string(networkBytes[8]) + " " + string(networkBytes[9]) + " " + string(networkBytes[10])
		peer1.Ports = string(networkBytes[11])
		peer1.Names = string(networkBytes[12])
		nodes = append(nodes, peer1)

		peer2.ID = string(networkBytes[13])
		peer2.Image = string(networkBytes[14])
		peer2.Status = string(networkBytes[21]) + " " + string(networkBytes[22]) + " " + string(networkBytes[23])
		peer2.Ports = string(networkBytes[24]) + " " + string(networkBytes[25])
		peer2.Names = string(networkBytes[26])
		nodes = append(nodes, peer2)

		orderer.ID = string(networkBytes[27])
		orderer.Image = string(networkBytes[28])
		orderer.Status = string(networkBytes[33]) + " " + string(networkBytes[34]) + " " + string(networkBytes[35])
		orderer.Ports = string(networkBytes[36])
		orderer.Names = string(networkBytes[37])
		nodes = append(nodes, orderer)
		return nodes
	} else {
		return nil
	}

}
