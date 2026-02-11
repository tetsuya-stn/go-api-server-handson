package api

import (
	"database/sql"
	"net/http"

	"github.com/tetsuya-stn/go-api-server-handson/controllers"
	"github.com/tetsuya-stn/go-api-server-handson/services"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	r := http.NewServeMux()

	r.HandleFunc("POST /article", aCon.PostArticleHandler)
	r.HandleFunc("GET /article/list", aCon.ArticleListHandler)
	r.HandleFunc("GET /article/{id}", aCon.ArticleDetailHandler)
	r.HandleFunc("POST /article/nice", aCon.PostNiceHandler)

	r.HandleFunc("POST /comment", cCon.PostCommentHandler)

	return r
}
