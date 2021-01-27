package unmarshal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

type Blockhash struct {
	Num       int
	Hash      string
	Prehash   string
	Blocktype int // 0 for config, 1 for transaction
}

type Channel struct {
	ChannelName string
	Bh          []Blockhash
}

func GetBlockNumberInfo(channelName string) (int, error) {
	// fetch block info
	command := `cd backend/scripts && ./showchannel.sh ` + channelName
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Output()
	// parse block info
	info, err := ioutil.ReadFile("./cache/" + channelName + "_blockinfo")
	if err != nil {
		return 0, err
	}
	fmt.Printf("[+] channel: %s|info: %s\n", channelName, info)
	jinfo := info[17:] // prefix(Blockchain info: )
	var f interface{}
	err = json.Unmarshal(jinfo, &f)
	if err != nil {
		return 0, err
	}
	num := f.(map[string]interface{})["height"].(float64)
	return int(num), nil
}

//Blocktype int // 0 for config, 1 for transaction

func GetBlockHash(channelName string, num int) (string, string, int, error) {
	// fetch block info
	command := `cd backend/scripts && ./fetchblock.sh ` + strconv.Itoa(num) + ` ` + channelName
	cmd := exec.Command("/bin/bash", "-c", command)
	// print(cmd)
	cmd.Output()

	//parse block info
	bcontent, err := ioutil.ReadFile("./cache/" + channelName + "_" + strconv.Itoa(num) + ".json")
	if err != nil {
		return "", "", 0, err
	}
	var f interface{}
	err = json.Unmarshal(bcontent, &f)
	if err != nil {
		return "", "", 0, err
	}

	// catch type convert panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	iheader := f.(map[string]interface{})
	header := iheader["header"].(map[string]interface{})

	var config int
	if num <= 2 {
		config = 0
	} else {
		config = 1
	}
	if header["previous_hash"] == nil {
		return header["data_hash"].(string), "NIL", config, nil
	} else {
		return header["data_hash"].(string), header["previous_hash"].(string), config, nil
	}
}

func GetChannelNameList() []string {
	//run channel list script
	command := `cd backend/scripts && ./getchannellist.sh`
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Output()

	clist, _ := ioutil.ReadFile("./cache/clist")

	// delete the first element and the last element
	if len(clist) == 0 {
		return nil
	} else {
		m := strings.Split(string(clist), "\n")
		channels := m[1 : len(m)-1]

		fmt.Println(channels)
		return channels
	}

}

func UnmarshalChannel() []Channel {
	// get channel list
	cnl := GetChannelNameList()

	if cnl != nil {
		cnln := len(cnl)
		channels := make([]Channel, cnln)

		for i, cn := range cnl {
			channels[i].ChannelName = cn

			blknum, err := GetBlockNumberInfo(cn)
			if err != nil {
				fmt.Printf("[-] Get block number fail: %s\n", err.Error())
				return nil
			}

			channels[i].Bh = make([]Blockhash, blknum)

			for j := 0; j < blknum; j++ {
				hash, prehash, config, err := GetBlockHash(cn, j)
				if err != nil {
					fmt.Printf("[-] Get block info fail: %s\n", err.Error())
				}

				channels[i].Bh[j].Num = j
				channels[i].Bh[j].Hash = hash
				channels[i].Bh[j].Prehash = prehash
				channels[i].Bh[j].Blocktype = config
			}
		}
		return channels

	} else {
		return nil
	}
}
