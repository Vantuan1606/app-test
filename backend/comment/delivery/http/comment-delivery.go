package http

import (
	"github.com/labstack/echo/v4"

	"github.com/Vantuan1606/app-test/domain"
)

type commentHTTPHandler struct {
	commentUsecase domain.ICommentUsecase
}

type responseErr struct {
	Error domain.Error `json:"error"`
}

type responseComment struct {
	Comment interface{} `json:"comment"`
}

type responseComments struct {
	Comments interface{} `json:"comments"`
}

func NewCommentHTTPHandler(e *echo.Echo, cm domain.ICommentUsecase) {
	handler := &commentHTTPHandler{
		commentUsecase: cm,
	}

	// e.GET("/hashtag", handler.Lists)
	// e.GET("/hashtag/:id", handler.GetHashtagDetail)

}
