package routes

import (
	handler "BlogAPI/handler"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	articleRouter := r.Group("/articles")
	{
		articleRouter.GET("/", handler.GetAllArticles)
		articleRouter.GET("/:id", handler.GetArticleByID)
		articleRouter.POST("/", handler.CreateManyPosts)
		articleRouter.DELETE("/", handler.DeleteArticles)
		articleRouter.PATCH("/:id", handler.UpdateContent)
		//articleRouter.POST("/", handlers.CreateManyArticle)
		//articleRouter.GET("/all", handlers.GetAllArticles)
		//articleRouter.GET("/all/topics", handlers.GetAllArticlesByTopic)
		//articleRouter.GET("/random", handlers.GetRandomArticles)
		//articleRouter.GET("/random/topics", handlers.GetRandomArticlesByTopic)
		//articleRouter.GET("/topics", handlers.GetAllTopic)
		//articleRouter.DELETE("/", handlers.DeleteArticles)
	}
}
