package testdata

import (
	"github.com/tetsuya-stn/go-api-server-handson/models"
)

var ArticleTestData = []models.Article{
	models.Article{
		Id:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "stn",
		NiceNum:  2,
	},
	models.Article{
		Id:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "stn",
		NiceNum:  4,
	},
}

var CommentTestData = []models.Comment{
	models.Comment{
		CommentId: 1,
		ArticleId: 1,
		Message:   "1st comment yeah",
	},
	models.Comment{
		CommentId: 2,
		ArticleId: 1,
		Message:   "welcome",
	},
}
