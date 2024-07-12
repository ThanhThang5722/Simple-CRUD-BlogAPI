package handler

import (
	"BlogAPI/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllArticles(ctx *gin.Context) {
	var a model.Article
	articles, err := a.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Fail to get all articles"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessResponse("Get all articles successfully", gin.H{
		"articles": articles,
	}))
}
func GetArticleByID(ctx *gin.Context) {
	var a model.Article
	err := a.GetByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Fail to get article by id"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessResponse("Get article successfully", gin.H{
		"articles": a,
	}))
}
