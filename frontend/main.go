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
	http.HandleFunc("/scripts.js", scriptsHandler)
	http.HandleFunc("/scripts.js.map", scriptsMapHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	fmt.Println("Listening on port 4321...")
	http.ListenAndServe(":4321", nil)
}

func scriptsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./frontend/scripts/scripts.js")
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func scriptsMapHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./frontend/scripts/scripts.js.map")
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/layout/favicon.ico")
}
