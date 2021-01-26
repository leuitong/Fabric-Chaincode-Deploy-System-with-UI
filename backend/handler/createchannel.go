package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	channelname := r.PostFormValue("channelname")
	//cn := r.Form[channelname][0]
	fmt.Printf("[+] fabric begin to create channel %s", channelname)
	command := `cd ~/go/src/github.com/hyperledger/fabric-samples/test-network && ./network.sh createChannel -c ` + channelname
	cmd := exec.Command("/bin/bash", "-c", command)

	//err := cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Printf(string(output))
	fmt.Printf("\n")
	fmt.Printf("[+] fabric channel %s create successful\n", channelname)
	t := template.Must(template.ParseFiles("frontend/main/createchannelsuccess.html"))
	t.Execute(w, err == nil)
}
