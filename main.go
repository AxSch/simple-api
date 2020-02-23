package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article model representation of dara
type Article struct {
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
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
		Article{
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content",
		},
		Article{
			Title:   "Hello 2",
			Desc:    "Article Description",
			Content: "Article Content",
		},
	}
	fmt.Println("Running at: 127.0.0.1:8000")
	handleRequests()
}
