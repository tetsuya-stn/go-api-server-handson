package testdata

import (
	"time"

	"github.com/tetsuya-stn/go-api-server-handson/models"
)

var ArticleTestData = []models.Article{
	models.Article{
		Id:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	},
	models.Article{
		Id:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum:  6,
	},
	models.Article{
		Id:       3,
		Title:    "3",
		Contents: "Third blog post",
		UserName: "saki",
		NiceNum:  6,
	},
	models.Article{
		Id:       4,
		Title:    "4th",
		Contents: "Forth blog post",
		UserName: "saki",
		NiceNum:  6,
	},
	models.Article{
		Id:       5,
		Title:    "5th",
		Contents: "Fifth blog post",
		UserName: "saki",
		NiceNum:  10,
	},
}

var CommentTestData = []models.Comment{
	models.Comment{
		CommentId: 1,
		ArticleId: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	},
	models.Comment{
		CommentId: 2,
		ArticleId: 2,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	},
}
