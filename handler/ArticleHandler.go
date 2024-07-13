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

/*
func CreateManyArticle(ctx *gin.Context) {
	var urls []string
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.ErrorResponse("Fail read to articles info 1"))
		return
	}

	err = json.Unmarshal(body, &urls)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.ErrorResponse("Fail to read articles info 2"))
		return
	}


	var a models.Article
	articles, _ := a.CreateMany(urls)

	ctx.JSON(http.StatusAccepted,
		models.SuccessResponse("Created articles successfully", gin.H{
			"articles": articles,
		}))
}
func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": data.Status,
				"error":  err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}
*/
