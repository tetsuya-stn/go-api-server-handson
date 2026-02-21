package services

import (
	"database/sql"
	"errors"

	"github.com/tetsuya-stn/go-api-server-handson/apperrors"
	"github.com/tetsuya-stn/go-api-server-handson/models"
	"github.com/tetsuya-stn/go-api-server-handson/repositories"
)

func (s *MyAppService) GetArticleService(articleId int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)

	go func(ch chan<- articleResult, db *sql.DB, articleId int) {
		article, err := repositories.SelectArticleDetail(db, articleId)
		ch <- articleResult{
			article: article,
			err:     err,
		}
	}(articleChan, s.db, articleId)

	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleId int) {
		commentList, err := repositories.SelectCommentList(db, articleId)
		ch <- commentResult{
			commentList: &commentList,
			err:         err,
		}
	}(commentChan, s.db, articleId)

	for range 2 {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(articleId int) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, articleId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return models.Article{}, err
	}

	resArticle, err := repositories.SelectArticleDetail(s.db, articleId)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	return resArticle, nil
}
