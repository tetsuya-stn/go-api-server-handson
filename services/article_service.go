package services

import (
	"github.com/tetsuya-stn/go-api-server-handson/models"
	"github.com/tetsuya-stn/go-api-server-handson/repositories"
)

func (s *MyAppService) GetArticleService(articleId int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleId)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleId)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(articleId int) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, articleId)
	if err != nil {
		return models.Article{}, err
	}

	resArticle, err := repositories.SelectArticleDetail(s.db, articleId)
	if err != nil {
		return models.Article{}, err
	}

	return resArticle, nil
}
