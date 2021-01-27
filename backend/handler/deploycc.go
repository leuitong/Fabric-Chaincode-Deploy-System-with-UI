package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func DeployChaincode(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	cn := r.PostFormValue("channelname")
	ccn := r.PostFormValue("chaincodename")
	lang := r.PostFormValue("language")
	v := r.PostFormValue("version")

	//cn := r.Form[channelname][0]
	//fn := r.Form[filename][0]

	fmt.Printf("[+] begin to deploy chaincode %s (%s,%s) in %s channel", ccn, lang, v, cn)
	t := template.Must(template.ParseFiles("frontend/main/deploysuccess.html"))
	command := `cd backend/scripts && ./dynamicdeploy.sh ` + cn + ` ` + ccn + ` ` + lang + ` ` + v
	cmd := exec.Command("/bin/bash", "-c", command)

	//err := cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Printf(string(output))
	fmt.Printf("\n")

	fmt.Printf("[+] deploy chaincode %s in %s successful.\n", ccn, cn)
	t.Execute(w, err == nil)
	//fmt.Fprintf(w, "[DeployChaincode] deploy chaincode fabcar in %s channel successful\n", cn)
}
