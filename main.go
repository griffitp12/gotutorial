package main

import (
	"fmt"
	"log"
	"net/http"

	/* "html/template" */
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// this doesn't work!

func serveCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	http.ServeFile(w, r, "static/styles.css")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// actual routed web endpoints

	myRouter.HandleFunc("/", AllDinos)
	myRouter.HandleFunc("/styles.css", serveCss)
	myRouter.HandleFunc("/adddino", AddDino)
	myRouter.HandleFunc("/deletedino", DeleteDinoPageServe)

	//only an api endpoint for now
	myRouter.HandleFunc("/api/v1/dino/{name}/{food}", UpdateDino).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	InitialMigration()
	handleRequests()
}
