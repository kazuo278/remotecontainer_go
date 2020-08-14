package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/kazuo278/remotecontainer_go/usecase"
)

// ArticleHandler : Article用のREST API
type ArticleHandler interface {
	GetArticlesHandler(w http.ResponseWriter, r *http.Request)
	PostArticleHandler(w http.ResponseWriter, r *http.Request)
}

type articleHandler struct {
	articleUsecase usecase.ArticleUsecase
}

// NewArticleHandler : Article用REST APIを返却する
func NewArticleHandler(usecase usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{
		articleUsecase: usecase,
	}
}

func (hundler *articleHandler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := hundler.articleUsecase.FindAll()
	errorCheck(err)
	json.NewEncoder(w).Encode(articles)
}

func (hundler *articleHandler) PostArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}
