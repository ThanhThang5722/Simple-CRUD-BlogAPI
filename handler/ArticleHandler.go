package handler

import (
	"BlogAPI/model"
	"BlogAPI/pkg/auth"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

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
	token := auth.GetTokenString(ctx)
	claims, err := auth.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized,
			model.ErrorResponse("Invalid token"))
		return
	}
	article_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var a model.Article_updating
	if err := ctx.ShouldBind(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	err = a.UpdateByID(claims.ID, article_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Fail to get article by id"))
		return
	}
	ctx.JSON(http.StatusOK, model.SuccessResponse("Update article's content successfully", gin.H{
		"status": "success",
	}))
}

/*
func GetUser(ctx *gin.Context) {
	token := auth.GetTokenString(ctx)
	claims, err := auth.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized,
			models.ErrorResponse("Invalid token"))
		return
	}

	user := models.User{}
	err = user.GetOne("_id", claims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.ErrorResponse("Fail to get user"))
		return
	}

	// Hide password
	user.Password = "*"

	ctx.JSON(http.StatusOK,
		models.SuccessResponse("Get user successfully", gin.H{"user": user}))
}
*/
