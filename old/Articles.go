package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	/* "gorm.io/driver/postgres"
	"gorm.io/gorm" */
	"io/ioutil"
	"net/http"
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

func ArticlesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articles")
	json.NewEncoder(w).Encode(Articles)
}

func SingleArticlePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Endpoint Hit: article", key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// parsing the path parameters
	vars := mux.Vars(r)
	// finding the id of the article we want to delete
	key := vars["id"]

	fmt.Println("Endpoint Hit: delete article", key)
	for index, article := range Articles {
		if article.Id == key {
			//update the articles array to remove the article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

	fmt.Fprintf(w, "Article %v Deleted", key)
}

func PutArticle(w http.ResponseWriter, r *http.Request) {
	// parsing the path parameters
	vars := mux.Vars(r)
	// finding the id of the article we want to delete
	key := vars["id"]

	// parse the new article data from the body of the request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle Article
	json.Unmarshal(reqBody, &newArticle)
	fmt.Fprintf(w, "%v", newArticle)

	// grab the index of the article that matches the given id
	var articleIndex int
	for index, article := range Articles {
		if article.Id == key {
			//update the articles array to remove the article
			articleIndex = index
		}
	}
	// update the Articles array with our newArticle
	Articles[articleIndex] = newArticle

	/* fmt.Fprintf(w, "Article %v Patched", key) */
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	//update our global Articles array to include our new article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}
