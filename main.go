package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handle(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}
	username := req.Form.Get("username")
	fmt.Fprintf(w, "Hello %s!", username)
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}
	id := req.Form.Get("id")
	username := ""
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	switch id_int {
	case 10:
		username = "Ten"
	case 8:
		username = "Eight"
	default:
		username = "None"
	}
	fmt.Fprintf(w, "Hello %s!", username)
}

func main() {

	http.HandleFunc("/", handle)
	http.HandleFunc("/user/get", handleGet)
	http.ListenAndServe(":8080", nil)

}
