package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func UploadChainCode(w http.ResponseWriter, r *http.Request) {
	ccn := r.PostFormValue("ccname")
	ccversion := r.PostFormValue("ccversion")
	language := r.PostFormValue("lang")
	rf, fh, err := r.FormFile("ccfile")
	t := template.Must(template.ParseFiles("frontend/main/uploadccsuccess.html"))

	if err != nil {
		fmt.Printf("[=] file upload fail: %s\n", err.Error())
		return
	}
	filen := fh.Filename
	fmt.Printf("[+] receive uploaded chaincode %s:v%s with %s.\n", ccn, ccversion, language)

	wf, err := os.OpenFile("chaincode/"+ccn+"."+language, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		fmt.Printf("[=] local open file fail.\n")
	}

	_, err = io.Copy(wf, rf)
	if err != nil {
		fmt.Printf("[-] copy file fail.\n")
	}

	fmt.Printf("[+] save chaincode file %s in server local.\n", filen)
	t.Execute(w, err == nil)

	//fmt.Fprintf(w, "[UploadChaincode] upload %s with language %s successful\n", ccn, language)

	//io.Copy(,f)
	//f.Read()
	//fmt.Printf("[+] upload %s with language %s successful\n", fn, language)

	//fmt.Fprintf(w, "[UploadChaincode] upload %s with language %s successful\n", fn, language)
}
