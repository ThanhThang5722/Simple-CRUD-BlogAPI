package handler

import (
	"BlogAPI/model"
	"encoding/json"
	"io"
	"log"
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
func CreateManyPosts(ctx *gin.Context) {
	var ListNewItems []model.Article_creation
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Fail to read"))
		return
	}

	err = json.Unmarshal(body, &ListNewItems)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Unmarshal Error"))
		return
	}
	var a model.Article
	if err = a.CreateItem(ListNewItems); err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Fail to Insert to DB"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessResponse("Get article successfully", gin.H{
		"Status": "Success",
	}))
}

func DeleteArticles(ctx *gin.Context) {
	var ListIdDel []int
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Fail to read"))
		return
	}

	err = json.Unmarshal(body, &ListIdDel)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Unmarshal Error"))
		return
	}
	var a model.Article
	if err = a.DeleteArticles(ListIdDel); err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Can't Delete"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessResponse("Deleted articles successfully", gin.H{
		"Status": "Success",
	}))
}

func UpdateContent(ctx *gin.Context) {
	var a model.Article_updating
	if err := ctx.ShouldBind(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	err := a.UpdateByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Fail to get article by id"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessResponse("Update article's content successfully", gin.H{
		"status": "success",
	}))
}
