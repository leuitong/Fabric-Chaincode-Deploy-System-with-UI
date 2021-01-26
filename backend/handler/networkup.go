package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func NetworkUp(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("[+] fabric network startup")
	command := `cd ~/go/src/github.com/hyperledger/fabric-samples/test-network && ./network.sh up`
	cmd := exec.Command("/bin/bash", "-c", command)

	//err := cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Println(string(output))
	//PrintSecondDot(8)

	fmt.Printf("\n")
	fmt.Printf("[+] fabric network startup successful\n")

	//fmt.Fprintf(w, "[NetworkStartup] fabric network startup Successful\n")
	t := template.Must(template.ParseFiles("frontend/main/networkupsuccess.html"))
	t.Execute(w, err == nil)
}
