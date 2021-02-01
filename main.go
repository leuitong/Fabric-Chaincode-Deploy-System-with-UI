package main

import (
	"contractdeploy/backend/handler"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func IndexHandeler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "")
}

func main() {
	//config
	PORT := "9090"

	fmt.Printf("[+] Dynamic Deploy network startup (DD-server)\n")
	//http.HandleFunc("/", sayhelloName) // 设置访问的路由
	http.Handle("/frontend/main/css/", http.StripPrefix("/frontend/main/css/", http.FileServer(http.Dir("frontend/main/css"))))
	http.Handle("/frontend/main/js/", http.StripPrefix("/frontend/main/js/", http.FileServer(http.Dir("frontend/main/js"))))
	http.Handle("/frontend/main/scss/", http.StripPrefix("/frontend/main/scss/", http.FileServer(http.Dir("frontend/main/scss"))))
	http.Handle("/frontend/assets/", http.StripPrefix("/frontend/assets/", http.FileServer(http.Dir("frontend/assets"))))
	http.Handle("/frontend/images/", http.StripPrefix("/frontend/images/", http.FileServer(http.Dir("frontend/images"))))
	http.Handle("/frontend/main/", http.StripPrefix("/frontend/main/", http.FileServer(http.Dir("frontend/main"))))

	http.HandleFunc("/", IndexHandeler)
	http.HandleFunc("/networkupsuccess", handler.NetworkUp)
	http.HandleFunc("/networkdownsuccess", handler.NetworkDown)
	http.HandleFunc("/createchannelsuccess", handler.CreateChannel)
	http.HandleFunc("/uploadsuccess", handler.UploadChainCode)
	http.HandleFunc("/deploysuccess", handler.DeployChaincode)
	http.HandleFunc("/shownetwork", handler.ShowNetwork)
	http.HandleFunc("/showchannel", handler.ShowChannel)
	http.HandleFunc("/showcontract", handler.ShowContract)
	http.HandleFunc("/querypromoney", handler.Querypromoney)
	http.HandleFunc("/querybuymoney", handler.Querybuymoney)
	http.HandleFunc("/querytransaction", handler.Querytransaction)
	err := http.ListenAndServe(":"+PORT, nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
