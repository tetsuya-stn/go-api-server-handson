package repositories_test

import (
	"testing"

	"github.com/tetsuya-stn/go-api-server-handson/models"
	"github.com/tetsuya-stn/go-api-server-handson/repositories"
	"github.com/tetsuya-stn/go-api-server-handson/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "sub test1",
			expected:  testdata.ArticleTestData[0],
		},
		{
			testTitle: "sub test2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.Id)
			if err != nil {
				t.Fatal(err)
			}

			if got.Id != test.expected.Id {
				t.Errorf("Id: got %d but want %d\n", got.Id, test.expected.Id)
			}

			if got.Title != test.expected.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: got %s but want %s\n", got.Contents, test.expected.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: got %s but want %s\n", got.UserName, test.expected.UserName)
			}

			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "test",
		UserName: "test",
	}

	expectedArticleNum := 15
	newArticle, err := repositories.InsertArticle(testDB, article)

	if err != nil {
		t.Error(err)
	}

	if newArticle.Id != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.Id)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from articles
		where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleId := 1
	var beforeNiceNum int
	var afterNiceNum int
	const sqlGetNice = ` select nice
        from articles
        where article_id = ?;
  `

	row := testDB.QueryRow(sqlGetNice, articleId)
	if err := row.Err(); err != nil {
		t.Fatal(err)
	}
	if err := row.Scan(&beforeNiceNum); err != nil {
		t.Fatal(err)
	}

	if err := repositories.UpdateNiceNum(testDB, articleId); err != nil {
		t.Fatal(err)
	}

	row = testDB.QueryRow(sqlGetNice, articleId)
	if err := row.Err(); err != nil {
		t.Fatal(err)
	}
	if err := row.Scan(&afterNiceNum); err != nil {
		t.Fatal(err)
	}

	if beforeNiceNum+1 != afterNiceNum {
		t.Errorf("NiceNum is expected %d but got %d\n", beforeNiceNum+1, afterNiceNum)
	}

	t.Cleanup(func() {
		const sqlStr = `
		update articles set nice = ? where article_id = ?
		`
		testDB.Exec(sqlStr, beforeNiceNum, articleId)
	})
}
