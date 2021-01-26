package handler

import (
	"contractdeploy/backend/unmarshal"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func ShowNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[+] fabric network display")
	command := `docker ps | grep 0.0.0.0:7051 && docker ps | grep 0.0.0.0:9051 && docker ps | grep orderer`
	cmd := exec.Command("/bin/bash", "-c", command)

	//err := cmd.Run()
	output, _ := cmd.Output()
	fmt.Println(output)

	//if err != nil {
	//	fmt.Println("Execute Command failed:" + err.Error())
	//	return
	//}

	outputnodes := unmarshal.UnmarshalNetwork(output)
	t := template.Must(template.ParseFiles("frontend/main/shownetwork.html"))
	t.Execute(w, outputnodes)
	//fmt.Fprintf(w, string(output))
}
