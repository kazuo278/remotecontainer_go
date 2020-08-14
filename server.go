package main

import (
	"log"
	"net/http"

	"github.com/kazuo278/remotecontainer_go/infrastructure/persistence"
	"github.com/kazuo278/remotecontainer_go/interface/handler"
	"github.com/kazuo278/remotecontainer_go/usecase"

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
