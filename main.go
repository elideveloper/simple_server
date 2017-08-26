package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var usersMap map[string]string

func handleGet(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}
	id := req.Form.Get("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "Hello %s!", usersMap[id])
}

func handleSet(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}
	id := req.Form.Get("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		return
	}
	username := req.Form.Get("username")
	usersMap[id] = username
	fmt.Fprintf(w, "Added %s!", usersMap[id])
}

func main() {
	usersMap = make(map[string]string)

	http.HandleFunc("/user/get", handleGet)
	http.HandleFunc("/user/set", handleSet)
	http.ListenAndServe(":8080", nil)
}
