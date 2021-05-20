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

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	
    myRouter.HandleFunc("/", HomePage)

	/* myRouter.HandleFunc("/", HomePage) */
	myRouter.HandleFunc("/dinos", AllDinos).Methods("GET")
	myRouter.HandleFunc("/dino/{name}/{food}", NewDino).Methods("POST")
    myRouter.HandleFunc("/dino/{name}", DeleteDino).Methods("DELETE")
    myRouter.HandleFunc("/dino/{name}/{food}", UpdateDino).Methods("PUT")

	/* myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", PutArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", SingleArticlePage) */

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	
	InitialMigration()
	/* Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	} */
	handleRequests()
}
