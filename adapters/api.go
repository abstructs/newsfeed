package adapters

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/abstructs/newsfeed/models"
	"github.com/abstructs/newsfeed/usecases"
	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func NewRouter(postAPI IPostAPI) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/posts", postAPI.CreatePost).Methods("POST")
	r.HandleFunc("/posts", postAPI.GetPosts).Methods("GET")
	return r
}

type IPostAPI interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
}

func NewPostAPI(logger *zap.Logger, usecase usecases.IPostUsecase) IPostAPI {
	return &postAPI{
		logger:  logger,
		usecase: usecase,
	}
}

type postAPI struct {
	logger  *zap.Logger
	usecase usecases.IPostUsecase
}

func (s *postAPI) CreatePost(w http.ResponseWriter, r *http.Request) {
	p := models.PostRequest{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPost, err := s.usecase.CreatePost(context.Background(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPost)
}

func (s *postAPI) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := s.usecase.GetPosts(context.Background())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
