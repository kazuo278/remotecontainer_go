package persistence

import (
	"github.com/kazuo278/remotecontainer_go/domain/model"
	"github.com/kazuo278/remotecontainer_go/domain/repository"

	// xorm
	"github.com/go-xorm/xorm"
	// Postgres用
	_ "github.com/lib/pq"
)

// ArticlePostgresRepository : ArticleRespositoryの実装 XORMを利用する
type articlePostgresRepository struct{}

// NewArticleRepository : ArticleRepositoryインターフェースの実装を返却する(DIP)
func NewArticleRepository() repository.ArticleRepository {
	return &articlePostgresRepository{}
}

// Persist : DBに１件の記事を登録する
func (repo *articlePostgresRepository) Persist(article *model.Article) error {
	engine, err := xorm.NewEngine("postgres", "dbname=tips host=localhost port=5432 user=admin password=password sslmode=disable")
	if err != nil {
		return err
	}
	defer engine.Close()
	// INSERT
	engine.Insert(&article)

	return nil
}

// FindAll : DBに登録された全ての記事を取得する
func (repo *articlePostgresRepository) FindAll() (*[]model.Article, error) {
	engine, err := xorm.NewEngine("postgres", "dbname=tips host=postgres port=5432 user=admin password=password sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer engine.Close()

	// SELECT
	var articles []model.Article
	engine.Find(&articles)

	return &articles, err
}
