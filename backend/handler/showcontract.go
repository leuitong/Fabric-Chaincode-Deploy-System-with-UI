package handler

import (
	"contractdeploy/backend/unmarshal"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func ShowContract(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[+] chaincode display")
	command := `cd backend/scripts && ./showcontract.sh`
	cmd := exec.Command("/bin/bash", "-c", command)

	//err := cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	outputcontracts := unmarshal.UnmarshalContract(output)
	t := template.Must(template.ParseFiles("frontend/main/showcontract.html"))
	t.Execute(w, outputcontracts)
	//fmt.Fprintf(w, string(output))
}
