package main

import (
	"log"
	"net/http"
	"tips/infrastructure/persistence"
	"tips/interface/handler"
	"tips/usecase"

	"github.com/gorilla/mux"
)

func main() {
	repository := persistence.NewArticleRepository()
	usecase := usecase.NewArticleUsecase(repository)
	handler := handler.NewArticleHandler(usecase)
	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.GetArticlesHandler).Methods("GET")
	router.HandleFunc("/article", handler.PostArticleHandler).Methods("POST")
	log.Println("Listen Server ...")
	log.Fatal(http.ListenAndServe(":9000", router))
}
