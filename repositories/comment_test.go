package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tetsuya-stn/go-api-server-handson/models"
	"github.com/tetsuya-stn/go-api-server-handson/repositories"
	"github.com/tetsuya-stn/go-api-server-handson/repositories/testdata"
)

func TestSelectCommentList(t *testing.T) {
	exptecedNum := len(testdata.CommentTestData)
	got, err := repositories.SelectCommentList(testDB, 1)

	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != exptecedNum {
		t.Errorf("want %d, but got %d comments\n", exptecedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleId: 1,
		Message:   "insertTest",
	}
	expectedCommentId := len(testdata.CommentTestData) + 1
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if newComment.CommentId != expectedCommentId {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentId, newComment.CommentId)
	}

	t.Cleanup(func() {
		const sqlStr = `
	delete from comments
	where message = ?
    `
		testDB.Exec(sqlStr, comment.Message)
	})
}
