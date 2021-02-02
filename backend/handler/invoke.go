package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func Invoke(w http.ResponseWriter, r *http.Request) {
	cn := r.PostFormValue("channelname")
	ctn := r.PostFormValue("chaincodename")
	m := r.PostFormValue("money")
	t := template.Must(template.ParseFiles("frontend/main/invokesuccess.html"))

	command := "cd backend/scripts && ./invoke.sh " + cn + " " + ctn + " " + m

	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	fmt.Println(string(output))

	fmt.Printf("\n")
	fmt.Printf("[+] transfer money successful\n")

	t.Execute(w, err == nil)
}
