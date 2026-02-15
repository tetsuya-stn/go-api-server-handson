package testdata

import (
	"github.com/tetsuya-stn/go-api-server-handson/models"
)

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) GetArticleService(articleId int) (models.Article, error) {
	return ArticleTestData[0], nil
}

func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return ArticleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return ArticleTestData, nil
}

func (s *serviceMock) PostNiceService(articleId int) (models.Article, error) {
	return ArticleTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return CommentTestData[0], nil
}
