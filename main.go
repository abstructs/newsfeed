package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type PostRequest struct {
	Text string `json:"text"`
}

type PostResponse struct {
	Text string
}

type Post struct {
	text string
}

var posts = []Post{}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// w.WriteHeader(http.StatusOK)
	p := PostRequest{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Text: {%+v}", p)

	newPost := Post{text: p.Text}

	posts = append(posts, newPost)
	json.NewEncoder(w).Encode(newPost)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

const serverPort string = "8000"

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	logger.Info(fmt.Sprintf("Server starting on port %s..", serverPort))

	r := mux.NewRouter()
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts", GetPosts).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", serverPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
