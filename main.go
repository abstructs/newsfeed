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
	Text string `json:"text"`
}

type Post struct {
	Text string `json:"text"`
}

type postService struct {
	logger *zap.Logger
	posts  []Post
}

func (s *postService) CreatePost(w http.ResponseWriter, r *http.Request) {
	p := PostRequest{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.logger.Info(fmt.Sprintf("Text: {%+v}", p))

	newPost := Post{Text: p.Text}

	s.posts = append(s.posts, newPost)
	json.NewEncoder(w).Encode(newPost)
}

func (s *postService) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s.logger.Sugar().Infof("posts: %+v", s.posts)
	json.NewEncoder(w).Encode(s.posts)
}

const serverPort string = "8000"

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	logger.Info(fmt.Sprintf("Server starting on port %s", serverPort))

	postService := postService{logger: logger, posts: []Post{}}

	r := mux.NewRouter()
	r.HandleFunc("/posts", postService.CreatePost).Methods("POST")
	r.HandleFunc("/posts", postService.GetPosts).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", serverPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
