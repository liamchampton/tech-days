package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./layout"))
	http.Handle("/", fs)
	http.HandleFunc("/scripts.js", scriptsHandler)
	http.HandleFunc("/scripts.js.map", scriptsMapHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Println("Listening on :4321...")
	log.Fatal(http.ListenAndServe(":4321", nil))
}

func scriptsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./scripts/scripts.js")
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func scriptsMapHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./scripts/scripts.js.map")
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./layout/favicon.ico")
}
