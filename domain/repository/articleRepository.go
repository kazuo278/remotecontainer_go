package repository

import (
	"github.com/kazuo278/remotecontainer_go/domain/model"
)

// ArticleRepository : Articleのリポジトリインターフェース
type ArticleRepository interface {
	Persist(*model.Article) error
	FindAll() (*[]model.Article, error)
}
