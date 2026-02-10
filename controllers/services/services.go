package services

import "github.com/tetsuya-stn/go-api-server-handson/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleId int) (models.Article, error)
	PostNiceService(articleId int) (models.Article, error)

	PostCommentService(comment models.Comment) (models.Comment, error)
}
