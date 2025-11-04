package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/tetsuya-stn/go-api-server-handson/models"
	"github.com/tetsuya-stn/go-api-server-handson/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}

	expected := models.Article{
		Id:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	}

	got, err := repositories.SelectArticleDetail(db, expected.Id)
	if err != nil {
		t.Fatal(err)
	}

	if got.Id != expected.Id {
		t.Errorf("Id: got %d but want %d\n", got.Id, expected.Id)
	}

	if got.Title != expected.Title {
		t.Errorf("Title: got %s but want %s\n", got.Title, expected.Title)
	}

	if got.Contents != expected.Contents {
		t.Errorf("Contents: got %s but want %s\n", got.Contents, expected.Contents)
	}

	if got.UserName != expected.UserName {
		t.Errorf("UserName: got %s but want %s\n", got.UserName, expected.UserName)
	}

	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}
