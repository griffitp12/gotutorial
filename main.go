package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func articlesPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: articles")
    json.NewEncoder(w).Encode(Articles)
}

func singleArticlePage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Println("Endpoint Hit: article", key)
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", articlesPage)
    myRouter.HandleFunc("/article/{id}", singleArticlePage)

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}