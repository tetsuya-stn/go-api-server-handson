package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tetsuya-stn/go-api-server-handson/apperrors"
	"github.com/tetsuya-stn/go-api-server-handson/controllers/services"
	"github.com/tetsuya-stn/go-api-server-handson/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		fmt.Printf("PostCommentService: %s", err.Error())
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
