package usecase

import (
	"github.com/kazuo278/remotecontainer_go/domain/model"
	"github.com/kazuo278/remotecontainer_go/domain/repository"
)

// ArticleUsecase : TIPSアプリのアプリケーション層インターフェース
type ArticleUsecase interface {
	FindAll() (*[]model.Article, error)
	Post(article *model.Article) error
}

type articleUsecase struct {
	articleRepository repository.ArticleRepository
}

// NewArticleUsecase : 記事に関するロジックを生成
func NewArticleUsecase(articleRepositoryImpl repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{
		articleRepository: articleRepositoryImpl,
	}
}

// FindAll : 全ての記事を取得する
func (uc *articleUsecase) FindAll() (*[]model.Article, error) {
	return uc.articleRepository.FindAll()
}

// Post : 記事を登録する
func (uc *articleUsecase) Post(article *model.Article) error {
	return uc.articleRepository.Persist(article)
}
