package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func NetworkDown(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[+] fabric network down")
	command := `cd ~/go/src/github.com/hyperledger/fabric-samples/test-network && ./network.sh down`
	cmd := exec.Command("/bin/bash", "-c", command)

	//err := cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Printf(string(output))
	//PrintSecondDot(8)

	fmt.Printf("\n")
	fmt.Printf("[+] fabric network down successful\n")

	//fmt.Fprintf(w, "[NetworkDown] fabric network down Successful\n")
	t := template.Must(template.ParseFiles("frontend/main/networkdownsuccess.html"))
	t.Execute(w, err == nil)
}
