package main

import (
	"fmt"
	"net/http"
	"strings"
	//"os"
)

func main() {
	//ip := os.Getenv("OPENSHIFT_GO_IP") + ":" + os.Getenv("OPENSHIFT_GO_PORT")
	ip := "localhost:8080"
	http.HandleFunc("/vasparun", respondHandler)
	http.ListenAndServe(ip, nil)
}

func respondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		text := r.FormValue("text")
		switch text {
			case "SelectAllUserTimes": 
		}
	}
}