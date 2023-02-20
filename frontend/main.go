package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/layout"))
	http.Handle("/", fs)
	http.HandleFunc("/scripts.js", scripts)
	http.HandleFunc("/scripts.js.map", scriptsMap)
	fmt.Println("Listening on port 4321...")
	http.ListenAndServe(":4321", nil)
}

func scripts(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./frontend/scripts/scripts.js")
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func scriptsMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./frontend/scripts/scripts.js.map")
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}
