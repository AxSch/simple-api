package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Article model representation of dara
type Article struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles -  a global Articles array
// simulates population of db
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get ya free homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	fmt.Fprintf(w, "%+v", string(reqBody))
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	key, err := strconv.Atoi(id)
	if err == nil {
		fmt.Println(key)
	}

	for index, article := range Articles {
		if article.ID == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf(string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	for index, artic := range Articles {
		if artic.ID == article.ID {
			Articles[index] = article
		}
	}
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keyString := params["id"]
	key, err := strconv.Atoi(keyString)
	if err == nil {
		fmt.Println(key)
	}

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func main() {
	Articles = []Article{
		Article{
			ID:      1,
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content",
		},
		Article{
			ID:      2,
			Title:   "Hello 2",
			Desc:    "Article Description",
			Content: "Article Content",
		},
	}
	fmt.Println("Running API at: 127.0.0.1:8000")
	handleRequests()
}
