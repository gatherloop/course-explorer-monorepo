package main

import (
	"course-explorer-monorepo/apps/api/content/core/module"
	"course-explorer-monorepo/apps/api/content/core/repository"
	"course-explorer-monorepo/apps/api/content/handler"
	"course-explorer-monorepo/libs/api/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){
  godotenv.Load()

	contentRepo := repository.NewContentRepository()
	contentUsecase := module.NewContentUsecase(contentRepo)
	contentHandler := handler.NewContentHandler(contentUsecase)

  mux := mux.NewRouter()
  mux.HandleFunc("/courses", contentHandler.GetContentList).Methods("GET")

	muxWithMiddlewares := middlewares.NewCorsMiddleware(mux.ServeHTTP)

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), muxWithMiddlewares))
}
