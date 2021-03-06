package handler

import (
	"contractdeploy/backend/unmarshal"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func Querypromoney(w http.ResponseWriter, r *http.Request) {
	cn := r.PostFormValue("channelname")
	ctn := r.PostFormValue("chaincodename")
	t := template.Must(template.ParseFiles("frontend/main/showprovider.html"))

	//tp := r.PostFormValue("querytype")
	//name := r.PostFormValue("name")
	/*
		if tp == "querypromoney" {

		} else if tp == "querybuymoney" {

		} else if tp == "querytransaction" {

		} else {
			fmt.Fprint(w, "fail with noknown querytype\n")
			return
		}*/
	//args := fmt.Sprintf("{\"Args\":[\"%s\", \"%s\", \"%s\"]}", tp, name)

	command := "cd backend/scripts && ./querypromoney.sh " + cn + " " + ctn

	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Println(string(output))
	outputprovider := unmarshal.UnmarshalProvider(output)

	fmt.Printf("\n")
	fmt.Printf("[+] query provider money successful\n")

	t.Execute(w, outputprovider)
}
