package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abstructs/newsfeed/adapters"
	"github.com/abstructs/newsfeed/usecases"
	"go.uber.org/zap"
)

const (
	serverPort     string = "8000"
	broker1Address string = "broker:9092"
)

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	logger.Info(fmt.Sprintf("Server starting on port %s", serverPort))

	writer := adapters.NewWriter(logger, []string{broker1Address})
	postUsecase := usecases.NewUsecase(logger, writer)
	postsAPI := adapters.NewPostAPI(logger, postUsecase)
	router := adapters.NewRouter(postsAPI)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", serverPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
