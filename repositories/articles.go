package repositories

import (
	"database/sql"

	"github.com/tetsuya-stn/go-api-server-handson/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
`
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	newArticle.Id = int(id)
	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
        select article_id, title, contents, username, nice
        from articles
        limit ? offset ?;
    `
	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Contents, &article.UserName, &article.NiceNum); err != nil {
			return nil, err
		}
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleId int) (models.Article, error) {
	const sqlStr = ` select *
        from articles
        where article_id = ?;
    `
	var article models.Article
	var createdTime sql.NullTime

	row := db.QueryRow(sqlStr, articleId)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	if err := row.Scan(&article.Id, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime); err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleId int) error {
	const sqlGetNice = ` select nice
        from articles
        where article_id = ?;
    `
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(sqlGetNice, articleId)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	var niceNum int
	if err := row.Scan(&niceNum); err != nil {
		return err
	}

	if _, err := tx.Exec(sqlUpdateNice, niceNum+1, articleId); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
