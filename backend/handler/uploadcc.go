package handler

import (
	"fmt"
	"net/http"
)

const (
	filename = "filename"
	lang     = "lang"
)

func UploadChainCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fn := r.Form[filename][0]
	lang := r.Form[lang][0]
	fmt.Printf("[+] receive uploaded chaincode %s.\n", fn)
	fmt.Printf("[+] upload %s with language %s successful\n", fn, lang)

	fmt.Fprintf(w, "[UploadChaincode] upload %s with language %s successful\n", fn, lang)
}
