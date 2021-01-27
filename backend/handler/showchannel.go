package handler

import (
	"contractdeploy/backend/unmarshal"
	"fmt"
	"html/template"
	"net/http"
)

func ShowChannel(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[+] channel display\n")
	t := template.Must(template.ParseFiles("frontend/main/showchannel.html"))
	channels := unmarshal.UnmarshalChannel()
	// render a template
	t.Execute(w, channels)
}
