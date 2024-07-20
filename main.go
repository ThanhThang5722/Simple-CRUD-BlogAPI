package main

import (
	"BlogAPI/pkg/auth"
	database "BlogAPI/pkg/database"
	"BlogAPI/pkg/middleware"
	"BlogAPI/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.ConnectMySQL()
	if err != nil {
		fmt.Println("Failt to connect MySQL")
		return
	}
	defer db.Close()

	auth.GenerateJWTKey()

	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ping successful",
			})
		})
	}
	//ROUTER DEFINE
	routes.ArticleRouter(api)
	routes.UserRouter(api)
	router.Run()
}
